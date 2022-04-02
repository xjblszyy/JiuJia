package main

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"jiujia/cmd"
	"jiujia/config"
	"jiujia/pkg/service"
)

var (
	cfgFile string
)

func RootCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use: "jiujia",
		RunE: func(command *cobra.Command, args []string) error {
			setups := []func() error{
				cmd.SetupPprofServer,
				cmd.SetupLogger,
				cmd.SetupHttpClient,
			}
			for _, setup := range setups {
				if err := setup(); err != nil {
					panic(err)
				}
			}

			s := service.New(*config.C, zap.L())
			if err := s.Start(); err != nil {
				panic(err)
			}
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			return nil
		},
	}

	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	return serverCmd
}

func main() {
	_ = RootCmd().Execute()
}
