package main

import (
	"fmt"
	"github.com/spf13/viper"
	"ngrokautomator/cmd"
)

func main() {
	// Initialize Viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	cmd.Execute()
}
