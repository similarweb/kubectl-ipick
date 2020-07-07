package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version of the release, the value injected by .goreleaser
	Version = `{{.Version}}`
	// Commit hash of the release, the value injected by .goreleaser
	Commit = `{{.Commit}}`
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
	fmt.Printf("Kubectl-interactive %s (%s)", Version, Commit)
	return nil
}
