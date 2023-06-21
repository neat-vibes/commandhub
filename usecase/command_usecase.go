package usecase

import (
	"context"
	"time"

	"github.com/neat-vibes/commandhub/entities"
)

type CommandUsecase struct {
	timeout          time.Duration
	commandProcessor entities.ICommandProcessor
}

func NewCommandUsecase(t time.Duration, c entities.ICommandProcessor) entities.ICommandUsecase {
	return &CommandUsecase{
		timeout:          t,
		commandProcessor: c,
	}
}

func (c *CommandUsecase) ExecuteCommand(ctx context.Context, req *entities.Request) (string, int, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	return c.commandProcessor.ExecuteCommand(ctx, req.Command, c.timeout)
}
