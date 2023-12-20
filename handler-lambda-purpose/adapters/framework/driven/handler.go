package driven

import (
  "context"

  "github.com/codeclout/AccountEd/pkg/monitoring"

  t "orchestrator/handler-types"
)

type Adapter struct {
  config  map[string]interface{}
  monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
  return &Adapter{
    config:  config,
    monitor: monitor,
  }
}

func (a *Adapter) WorkWithDependencies(ctx context.Context, in *t.HandlerRequest) (*t.HandlerResponse, *t.CustomError) {
  return nil, &t.CustomError{Message: "not implemented"}
}
