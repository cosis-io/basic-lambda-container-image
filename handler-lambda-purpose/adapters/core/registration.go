package core

import (
  "context"

  "github.com/codeclout/AccountEd/pkg/monitoring"

  t "orchestrator/registration-types"
)

type Adapter struct {
  config  map[string]interface{}
  monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
  return &Adapter{config: config, monitor: monitor}
}

func (a *Adapter) ProcessTenant(ctx context.Context, in t.CreateTenantResponse) (t.CoreTenantResponse, *t.CustomError) {
  return t.CoreTenantResponse{}, nil
}

func (a *Adapter) BusinessLogicWork(ctx context.Context, in t.CreateTenantAdminResponse) (t.CoreTenantAdminResponse, *t.CustomError) {
  return t.CoreTenantAdminResponse{}, nil
}

func (a *Adapter) ProcessTenantProvisioning(ctx context.Context, in t.ProvisionTenantResponse) (t.CoreProvisionTenantResponse, *t.CustomError) {
  return t.CoreProvisionTenantResponse{}, nil
}

func (a *Adapter) CoreTenantAggregateResponse(ctx context.Context, in t.CoreTenantAggregate) (t.RegisterTenantResponse, *t.CustomError) {
  return t.RegisterTenantResponse{}, nil
}
