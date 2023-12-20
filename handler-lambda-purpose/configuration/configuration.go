package configuration

import (
  "errors"
  "fmt"
  "os"
  "path/filepath"
  "reflect"

  "github.com/codeclout/AccountEd/pkg/monitoring"
  "github.com/hashicorp/hcl/v2"
  "github.com/hashicorp/hcl/v2/hclsimple"
)

type environment struct {
  AWSRegion          string
  RuntimeEnvironment string
}

type metadataAndSettings struct {
  Metadata Metadata `hcl:"Metadata,block"`
  Settings Settings `hcl:"Settings,block"`
}

type Metadata struct {
  ServiceDescription string `hcl:"description"`
  ServiceName        string `hcl:"service"`
  ServiceVersion     string `hcl:"version"`
}

type Settings struct {
  ApiVersion   string  `hcl:"api_version"`
  Charset      string  `hcl:"charset"`
  HttpTimeout  float64 `hcl:"http_timeout"`
  IsAppGetOnly bool    `hcl:"is_app_get_only"`
}

type Adapter struct {
  monitor monitoring.Adapter
}

func NewAdapter(monitor monitoring.Adapter) *Adapter {
  return &Adapter{
    monitor: monitor,
  }
}

func (a *Adapter) LoadConfiguration() *map[string]interface{} {
  var metadataAndSettings metadataAndSettings
  var out = make(map[string]interface{})
  var s string

  currentLocation, _ := os.Getwd()
  fileLocation := filepath.Join(currentLocation, os.Getenv("STATIC_CONFIGURATION_PATH"))

  e := hclsimple.DecodeFile(fileLocation, nil, &metadataAndSettings)
  if e != nil {
    var x hcl.Diagnostics
    if errors.As(e, &x) {
      for _, x := range x {
        if x.Severity == hcl.DiagError {
          a.monitor.LogGenericError(fmt.Sprintf("Failed to load registration runtime staticConfig: %s", x))
        }
      }
      panic(e)
    }
  }

  env := environment{
    AWSRegion:          os.Getenv("AWS_REGION"),
    RuntimeEnvironment: os.Getenv("ENVIRONMENT"),
  }

  runtimeEnv := reflect.ValueOf(&env).Elem()
  for i := 0; i < runtimeEnv.NumField(); i++ {
    out[runtimeEnv.Type().Field(i).Name] = runtimeEnv.Field(i).Interface()
  }

  metadata := reflect.ValueOf(&metadataAndSettings.Metadata).Elem()
  for i := 0; i < metadata.NumField(); i++ {
    out[metadata.Type().Field(i).Name] = metadata.Field(i).Interface()
  }

  settings := reflect.ValueOf(&metadataAndSettings.Settings).Elem()
  for i := 0; i < settings.NumField(); i++ {
    out[settings.Type().Field(i).Name] = settings.Field(i).Interface()
  }

  for k, v := range out {
    switch x := v.(type) {
    case string:
      if x == (s) {
        a.monitor.LogGenericError(fmt.Sprintf("Registration Orchestrator:%s is not defined in the environment", k))
        os.Exit(1)
      }
    case bool:
      continue
    case float64:
      continue
    default:
      panic(errors.New("invalid configuration type"))
    }
  }

  return &out
}
