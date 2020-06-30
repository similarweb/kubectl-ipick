package command

import (
	"bytes"
	"fmt"
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
	}).Error("execute command")

	cmd := exec.Command(command, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr
	sshConnectError := cmd.Run()

	if sshConnectError != nil {
		_, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		log.WithFields(log.Fields{
			"command": command,
			"args":    args,
			"error":   errStr,
		}).Warn(fmt.Sprintf("command failed"))
	}

	return sshConnectError

}
