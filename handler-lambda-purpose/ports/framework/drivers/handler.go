package drivers

import (
  "context"

  t "orchestrator/handler-types"
)

type HandlerDriver interface {
  Handler(ctx context.Context, in *t.HandlerRequest) (*t.HandlerResponse, *t.CustomError)
}
