package cmd

import (
	"os"

	"github.com/approvers/qip/pkg/server/router"
	"github.com/approvers/qip/pkg/utils/config"
	"github.com/spf13/cobra"
)

func serve() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "start Qip Application Server",
		Run: func(cmd *cobra.Command, args []string) {
			f, err := os.Open("config.yml")
			if err != nil {
				panic("failed to read config file")
			}
			_, err = config.LoadConfig(f)
			if err != nil {
				panic("failed to load config")
			}
			router.StartServer(7000)
		},
	}
}
