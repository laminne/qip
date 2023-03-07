package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "qip",
	Short: "Qip - Minimum ActivityPub Implementation",
}

// Start Command Entrypoint
func Start() {
	root.AddCommand(serve())
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
