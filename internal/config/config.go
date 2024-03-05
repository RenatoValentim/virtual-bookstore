package config

import (
	"log"

	"github.com/RenatoValentim/virtual-bookstore/internal/constants/environments"
	"github.com/spf13/viper"
)

func setDefaults() {
	viper.SetDefault(environments.Environment, environments.Dev)
	viper.SetDefault(environments.ServerHost, "localhost")
	viper.SetDefault(environments.ServerPort, 8000)
	viper.SetDefault(environments.DBHost, "localhost")
	viper.SetDefault(environments.DBPort, 5432)
	viper.SetDefault(environments.DBUser, "postgres")
	viper.SetDefault(environments.DBPassword, 1234)
}

func bindEnvironmentVariables() {
	// INFO: Env Prefix
	viper.SetEnvPrefix("VIRTUALBOOKSTORE")

	// INFO: Server Variables
	viper.BindEnv(environments.ServerHost)
	viper.BindEnv(environments.ServerPort)

	// INFO: Databse Variables
	viper.BindEnv(environments.DBHost)
	viper.BindEnv(environments.DBPort)
	viper.BindEnv(environments.DBName)
	viper.BindEnv(environments.DBUser)
	viper.BindEnv(environments.DBPassword)

	// INFO: Running on environment
	viper.BindEnv(environments.Environment)
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
