package utils

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/nakachan-ing/reflect-cli/model"
)

func OpenEditor(filePath string, config model.Config) error {
	time.Sleep(2 * time.Second)
	c := exec.Command(config.Editor, filePath)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return fmt.Errorf("failed to open editor (%s): %w", filePath, err)
	}
	return nil
}
