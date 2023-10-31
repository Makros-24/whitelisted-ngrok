package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"ngrokautomator/pkg/ngrokautomator"
)

var run = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Usage: %s <address:port>", args[0])
		}

		whitelist, err := cmd.Flags().GetStringArray("whitelist")
		if err != nil {
			panic(err)
		}
		if err := ngrokautomator.Run(context.Background(), args[0], whitelist); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	run.Flags().StringArrayP("whitelist", "w", []string{}, "Whitelist for connecting IP addresses sepperated by comma ','")
	_ = run.MarkFlagRequired("whitelist")
	rootCmd.AddCommand(run)
}
