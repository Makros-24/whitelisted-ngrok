package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"golang.ngrok.com/ngrok/config"
	"log"
	"ngrokautomator/pkg/ngrokautomator"
)

var tcpTunnel = &cobra.Command{
	Use: "tcp",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Usage: %s <address:port>", args[0])
		}

		whitelist, err := cmd.Flags().GetStringArray("whitelist")
		if err != nil {
			panic(err)
		}
		if err := ngrokautomator.Run(context.Background(), args[0], whitelist, config.TCPEndpoint()); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	tcpTunnel.Flags().StringArrayP("whitelist", "w", []string{}, "Whitelist for connecting IP addresses")
	_ = tcpTunnel.MarkFlagRequired("whitelist")
	rootCmd.AddCommand(tcpTunnel)
}
