package command

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {

	t.Run("valid", func(t *testing.T) {

		pwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("cold not get working dir")
		}

		now := time.Now().Unix()
		filePath := fmt.Sprintf("%s/testutils/%d.txt", pwd, now)
		Run("touch", []string{filePath})

		_, err = os.Open(filePath)
		if err != nil {
			t.Fatalf("ssh command not success. error: %v", err)
		}

		Run("rm", []string{"-f", filePath})
		_, err = os.Open(filePath)
		if err == nil {
			t.Fatalf("ssh command should return error. error: %v", err)
		}

	})
	t.Run("invalid command", func(t *testing.T) {
		pwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("cold not get working dir")
		}

		now := time.Now().Unix()
		filePath := fmt.Sprintf("%s/testutils/%d.txt", pwd, now)
		Run("error-command", []string{filePath})

		_, err = os.Open(filePath)
		if err == nil {
			t.Fatalf("ssh command not success. error: %v", err)
		}

	})
}
