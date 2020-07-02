package command

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Run will execute commands
func Run(command string, args []string) error {

	log.WithFields(log.Fields{
		"command": command,
		"args":    strings.Join(args, " "),
	}).Debug("execute command")

	cmd := exec.Command(command, args...)
	var stderr bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {

		errStr := stderr.String()
		log.WithFields(log.Fields{
			"command": command,
			"args":    args,
		}).Info(errStr)

		log.Error(errStr)
	}

	return err

}
