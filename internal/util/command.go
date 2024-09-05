package util

import (
	"bytes"
	"context"
	"os/exec"

	"gorm.io/gorm/logger"
)

func Execute(command string) error {
	logger.Default.Error(context.Background(), "execute: %s", command)

	var outBuffer, errBuffer bytes.Buffer
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	err := cmd.Run()

	logger.Default.Info(context.Background(), "execute [%s]: %s", command, outBuffer.String())
	if err != nil {
		logger.Default.Error(context.Background(), "execute [%s]: %s", command, errBuffer.String())
	}

	return err
}
