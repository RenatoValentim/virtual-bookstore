package config

import (
	"log"

	"github.com/RenatoValentim/virtual-bookstore/internal/constants/environments"
	"github.com/spf13/viper"
)

func setDefaults() {
	viper.SetDefault("environment", environments.Dev)
	viper.SetDefault("server_host", "localhost")
	viper.SetDefault("server_port", 8000)
	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_port", 5432)
	viper.SetDefault("db_user", "postgres")
	viper.SetDefault("db_password", 1234)
}

func bindEnvironmentVariables() {
	// INFO: Env Prefix
	viper.SetEnvPrefix("VIRTUALBOOKSTORE")

	// INFO: Server Variables
	viper.BindEnv("server_host")
	viper.BindEnv("server_port")

	// INFO: Databse Variables
	viper.BindEnv("db_host")
	viper.BindEnv("db_port")
	viper.BindEnv("db_name")
	viper.BindEnv("db_user")
	viper.BindEnv("db_password")

	// INFO: Running on environment
	viper.BindEnv("environment")
}

func LoadConfig() {
	setDefaults()

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			bindEnvironmentVariables()
		} else {
			log.Fatalf("Failed to load configuration from configuration file. %s", err.Error())
		}
	}
}
