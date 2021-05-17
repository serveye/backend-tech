package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ViperConfig(appConfig *AppConfig) {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "dev_db")
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	// Set the filename of the configuration file
	viper.SetConfigName("config")

	// Set the path to look for the configuration file
	viper.AddConfigPath(".")

	// enable viper to read env variables
	viper.AutomaticEnv()

	// set config file type
	viper.SetConfigType("yml")
}


