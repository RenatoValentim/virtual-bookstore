package config

import "github.com/spf13/viper"

func setDefaults() {
	viper.SetDefault("environment", "dev")
}