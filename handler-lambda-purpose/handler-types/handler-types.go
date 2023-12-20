package handler_types

import (
  "github.com/codeclout/AccountEd/pkg/monitoring"
)

var monitor = monitoring.NewAdapter()

type Enum int

const (
  A Enum = iota + 1
  B
  C
  D
)

const (
  STAGENAME = "sch00l.lambda.handler"
)

type CustomError struct {
  Message string
}

func (e CustomError) Error() string {
  monitor.LogGenericError(e.Message)
  return e.Message
}

type HandlerRequest struct {
}

type HandlerResponse struct {
}
