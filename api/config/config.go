package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string `json:"name"`
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"app"`
}

func GetConfig() Config {
	const (
		configName string = "config"
		configType string = "yml"
	)
	var config Config

	baseDirectory, error := os.Getwd()
	if error != nil {
	}

	viper.AddConfigPath(baseDirectory)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	
	error = viper.ReadInConfig()
	if error != nil {
	}

	error = viper.Unmarshal(&config)
	if error != nil {
		
	}

	return config
}

