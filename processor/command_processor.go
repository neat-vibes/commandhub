package processor

import (
	"bytes"
	"context"
	"os/exec"
	"syscall"
	"time"
)

type CommandProcessor struct {
}

func NewCommandProcessor() *CommandProcessor {
	return &CommandProcessor{}
}

func (c *CommandProcessor) ExecuteCommand(ctx context.Context, cmd string, timeout time.Duration) (string, int, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmdWithContext := exec.CommandContext(ctx, "sh", "-c", cmd)

	outputBuf := bytes.NewBuffer(nil)
	cmdWithContext.Stdout = outputBuf
	cmdWithContext.Stderr = outputBuf

	cmdWithContext.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true, // placing the child process in a new process group
	}

	err := cmdWithContext.Start()
	if err != nil {
		return "", 0, err
	}

	done := make(chan error)
	go func() {
		done <- cmdWithContext.Wait()
	}()

	select {
	case <-ctx.Done():
		// timeout or context cancellation
		syscall.Kill(-cmdWithContext.Process.Pid, syscall.SIGTERM) // Terminate all processes in the process group

		return "", 124, ctx.Err()
	case err := <-done:
		// command execution completed
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				exitCode := exitError.ExitCode()
				return outputBuf.String(), exitCode, nil
			}
			return "", 0, err
		}
		return outputBuf.String(), 0, nil
	}
}
