package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	// version of the release, the value injected by .goreleaser
	version = `{{.Version}}`
	// commit hash of the release, the value injected by .goreleaser
	commit = `{{.Commit}}`
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the kubectl-interactive version",
	Args:  cobra.NoArgs,
	RunE:  runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, _ []string) error {
	fmt.Printf("Kubectl-interactive %s (%s)", version, commit)
	return nil
}
