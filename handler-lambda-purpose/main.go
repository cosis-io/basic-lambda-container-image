package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/codeclout/AccountEd/pkg/monitoring"

  apiAdapter "orchestrator/adapters/api"
  coreAdapter "orchestrator/adapters/core"
  drivenAdapter "orchestrator/adapters/framework/driven"
  "orchestrator/adapters/framework/drivers"
  "orchestrator/configuration"
)

var monitor *monitoring.Adapter
var configurationAdapter *configuration.Adapter

func init() {
  monitor = monitoring.NewAdapter()
  configurationAdapter = configuration.NewAdapter(*monitor)

  // cfg, e := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
  // if e != nil {
  //   panic(t.CustomError{
  //     Message: fmt.Sprintf("failed loading config, %v", e),
  //   })
  // }
  //
  // eventBridgeClient = eventbridge.NewFromConfig(cfg)
}

func main() {
  runtimeConfig := configurationAdapter.LoadConfiguration()

  core := coreAdapter.NewAdapter(*runtimeConfig, *monitor)
  driven := drivenAdapter.NewAdapter(*runtimeConfig, *monitor)
  api := apiAdapter.NewAdapter(*runtimeConfig, *core, *driven, *monitor)
  driver := drivers.NewAdapter(*runtimeConfig, api, *monitor)

  lambda.Start(driver.Handler)
}
