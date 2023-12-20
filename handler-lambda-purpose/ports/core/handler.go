package core

import (
  "context"

  t "orchestrator/handler-types"
)

type HandlerCore interface {
  BusinessLogicWork(ctx context.Context, in *t.HandlerResponse) (*t.HandlerResponse, *t.CustomError)
}
