package core

import (
	"context"

	t "orchestrator/registration-types"
)

type RegistrationCore interface {
	ProcessTenant(ctx context.Context, in t.CreateTenantResponse) (t.CoreTenantResponse, *t.CustomError)
	ProcessTenantAdmin(ctx context.Context, in t.CreateTenantAdminResponse) (t.CoreTenantAdminResponse, *t.CustomError)
	ProcessTenantProvisioning(ctx context.Context, in t.ProvisionTenantResponse) (t.CoreProvisionTenantResponse, *t.CustomError)
	CoreTenantAggregateResponse(ctx context.Context, in t.CoreTenantAggregate) (t.RegisterTenantResponse, *t.CustomError)
}
