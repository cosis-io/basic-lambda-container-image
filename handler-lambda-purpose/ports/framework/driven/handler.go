package driven

import (
  "context"

  t "orchestrator/handler-types"
)

type HandlerDriven interface {
  WorkWithDependencies(ctx context.Context, registration *t.HandlerRequest) (*t.HandlerResponse, *t.CustomError)
}
