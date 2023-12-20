package drivers

import (
  "context"
  "fmt"

  "github.com/aws/aws-lambda-go/lambdacontext"
  "github.com/codeclout/AccountEd/pkg/monitoring"

  t "orchestrator/handler-types"
  "orchestrator/ports/api"
)

type LambdaName string

type Adapter struct {
  api     api.HandlerAPI
  config  map[string]interface{}
  monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, api api.HandlerAPI, monitor monitoring.Adapter) *Adapter {
  return &Adapter{
    api:     api,
    config:  config,
    monitor: monitor,
  }
}

func (a *Adapter) setContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
  deadline, _ := ctx.Deadline()

  name, ok := a.config["ServiceName"].(string)
  if !ok {
    panic(t.CustomError{
      Message: "ServiceName not available in environment",
    })
  }

  n := LambdaName(name)

  ctx = context.WithValue(ctx, n, lambdacontext.FunctionName)
  ctx, cancel := context.WithDeadline(ctx, deadline)

  return ctx, cancel
}

func (a *Adapter) processHandler(in *t.HandlerRequest) (*t.HandlerRequest, *t.CustomError) {

  if in == nil {
    return nil, &t.CustomError{Message: "registration payload is invalid"}
  }

  // handle StageName here
  return in, nil
}

func (a *Adapter) Handler(ctx context.Context, in *t.HandlerRequest) (*t.HandlerResponse, *t.CustomError) {
  metadata, e := a.processHandler(in)
  if e != nil {
    panic(&t.CustomError{Message: e.Error()})
  }

  ch := make(chan *t.HandlerResponse, 1)
  ech := make(chan *t.CustomError, 1)

  ctx, cancel := a.setContextTimeout(ctx)

  defer cancel()
  a.api.Handler(ctx, metadata, ch, ech)

  select {
  case <-ctx.Done():
    panic(t.CustomError{
      Message: "the operation has been timed out",
    })

  case out := <-ch:
    a.monitor.LogGenericInfo(fmt.Sprintf("Successfully created %s", out))
    return out, nil

  case e := <-ech:
    panic(t.CustomError{
      Message: e.Error(),
    })
  }
}
