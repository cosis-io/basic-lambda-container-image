package api

import (
  "context"

  "github.com/codeclout/AccountEd/pkg/monitoring"

  coreAdapter "orchestrator/adapters/core"
  drivenAdapter "orchestrator/adapters/framework/driven"
  t "orchestrator/handler-types"
)

type Adapter struct {
  config  map[string]interface{}
  core    coreAdapter.Adapter
  driven  drivenAdapter.Adapter
  monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, core coreAdapter.Adapter, driven drivenAdapter.Adapter, monitor monitoring.Adapter) *Adapter {
  return &Adapter{
    config:  config,
    core:    core,
    driven:  driven,
    monitor: monitor,
  }
}

func (a *Adapter) Handler(ctx context.Context, in *t.HandlerRequest, ch chan *t.HandlerResponse, ech chan *t.CustomError) {
  driven, e := a.driven.WorkWithDependencies(ctx, in)
  if e != nil {
    ech <- e
    return
  }

  core, e := a.core.BusinessLogicWork(ctx, driven)
  if e != nil {
    ech <- e
    return
  }

  ch <- &core
  return
}
