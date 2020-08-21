package config

import (
	"errors"

	"github.com/spf13/viper"
)

// Configuration holds all the config related to this application
type Configuration struct {
	ApplicationPort string
	DBType          string
	DBUserName      string
	DBPassword      string
	DBName          string
	DBHost          string
	DBPort          string
}

// GetConfig return a abject that contains the Configuration
func GetConfig() (config Configuration, err error) {
	viper.SetConfigFile(`config.yaml`)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	config.ApplicationPort = viper.GetString("APPLICATION_PORT")
	if len(config.ApplicationPort) == 0 {
		err = errors.New("Application Port is not provided")
	}
	config.DBType = viper.GetString("DB_TYPE")
	if len(config.DBType) == 0 {
		err = errors.New("DB Type is not provided")
	}
	config.DBName = viper.GetString("DB_NAME")
	if len(config.DBName) == 0 {
		err = errors.New("DB Name is not provided")
	}

	config.DBUserName = viper.GetString("DB_USERNAME")
	if len(config.DBUserName) == 0 {
		err = errors.New("DB Username is not provided")
	}
	config.DBPassword = viper.GetString("DB_PASSWORD")
	// if len(config.ApplicationPort) == 0 {
	// 	err = errors.New("DB Password is not provided")
	// }
	config.DBHost = viper.GetString("DB_HOST")
	if len(config.DBHost) == 0 {
		err = errors.New("DB Host is not provided")
	}
	config.DBPort = viper.GetString("DB_PORT")
	if len(config.DBPort) == 0 {
		err = errors.New("DB Port is not provided")
	}
	return config, err
}
