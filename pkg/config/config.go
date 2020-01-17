package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Load the configuration file
func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./configs")
	viper.AddConfigPath("/etc/4okr/")
	viper.AddConfigPath("$HOME/.4okr")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file not found: %s", err))
		} else {
			panic(fmt.Errorf("config file was found but another error was produced: %s", err))
		}
	}
}
