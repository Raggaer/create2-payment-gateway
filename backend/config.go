package main

import (
	"github.com/spf13/viper"
)

// LoadConfig loads the configuration
func LoadConfig(path string) error {
	viper.SetConfigName(".env")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.WatchConfig()
	return viper.ReadInConfig()
}
