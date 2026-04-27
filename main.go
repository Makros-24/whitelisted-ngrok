package main

import (
	"log"

	"github.com/spf13/viper"
	"ngrokautomator/cmd"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error reading config file: %v", err)
	}
	cmd.Execute()
}
