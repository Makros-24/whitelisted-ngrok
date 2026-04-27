package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"golang.ngrok.com/ngrok/config"
	"ngrokautomator/pkg/ngrokautomator"
)

var httpTunnel = &cobra.Command{
	Use:   "http <address:port>",
	Short: "Create an HTTP tunnel to the given destination",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		whitelist, err := cmd.Flags().GetStringArray("whitelist")
		if err != nil {
			log.Fatal(err)
		}
		if err := ngrokautomator.Run(context.Background(), args[0], whitelist, config.HTTPEndpoint()); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	httpTunnel.Flags().StringArrayP("whitelist", "w", []string{}, "IP addresses allowed to connect (repeatable)")
	_ = httpTunnel.MarkFlagRequired("whitelist")
	rootCmd.AddCommand(httpTunnel)
}
