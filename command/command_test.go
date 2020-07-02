package command

import (
	"testing"
)

func TestCommand(t *testing.T) {

	t.Run("valid", func(t *testing.T) {

		err := Run("ls", []string{})

		if err != nil {
			t.Fatalf("command failed to execute. error: %v", err)
		}

	})
	t.Run("invalid command", func(t *testing.T) {

		err := Run("error-command", []string{})

		if err == nil {
			t.Fatalf("command failed to execute. error: %v", err)
		}

	})
}
