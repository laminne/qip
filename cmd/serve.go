package cmd

import (
	"github.com/approvers/qip/pkg/server/router"
	"github.com/spf13/cobra"
)

func serve() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "start Qip Application Server",
		Run: func(cmd *cobra.Command, args []string) {
			router.StartServer(7000)
		},
	}
}
