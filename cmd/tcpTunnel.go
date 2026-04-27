package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"golang.ngrok.com/ngrok/config"
	"ngrokautomator/pkg/ngrokautomator"
)

var tcpTunnel = &cobra.Command{
	Use:   "tcp <address:port>",
	Short: "Create a TCP tunnel to the given destination",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		whitelist, err := cmd.Flags().GetStringArray("whitelist")
		if err != nil {
			log.Fatal(err)
		}
		if err := ngrokautomator.Run(context.Background(), args[0], whitelist, config.TCPEndpoint()); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	tcpTunnel.Flags().StringArrayP("whitelist", "w", []string{}, "IP addresses allowed to connect (repeatable)")
	_ = tcpTunnel.MarkFlagRequired("whitelist")
	rootCmd.AddCommand(tcpTunnel)
}
