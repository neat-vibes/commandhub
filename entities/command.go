package entities

import (
	"context"
	"encoding/json"
	"time"
)

type Request struct {
	Command string `json:"command"`
}

type Response struct {
	Result   json.RawMessage `json:"result"`
	ExitCode int             `json:"exitcode"`
	Err      string          `json:"error"`
}

type ICommandProcessor interface {
	ExecuteCommand(c context.Context, cmd string, timeout time.Duration) (string, int, error)
}

type ICommandUsecase interface {
	ExecuteCommand(c context.Context, req *Request) (string, int, error)
}
