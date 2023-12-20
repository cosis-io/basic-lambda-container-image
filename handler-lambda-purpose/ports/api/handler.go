package api

import (
  "context"

  t "orchestrator/handler-types"
)

type HandlerAPI interface {
  Handler(ctx context.Context, in *t.HandlerRequest, ch chan *t.HandlerResponse, ech chan *t.CustomError)
}
