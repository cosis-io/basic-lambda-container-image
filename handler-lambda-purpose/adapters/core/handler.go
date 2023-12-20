package core

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
  return &Adapter{config: config, monitor: monitor}
}

func (a *Adapter) BusinessLogicWork(ctx context.Context, in *t.HandlerResponse) (*t.HandlerResponse, *t.CustomError) {
  return nil, &t.CustomError{Message: "Not Implemented"}
}
