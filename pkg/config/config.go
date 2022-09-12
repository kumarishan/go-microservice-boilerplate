package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string `default:"8080"`
	Environment string `default:"dev"`

	Logger struct {
		LogLevel string `default:"info"`
		FileName string `default:"server.log"`
	}

	DB struct {
		MySql struct {
			Database string `default:"catalog"`
		}
	}
}

func NewConfig() (*Config, error) {
	config := Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	fmt.Println(config)

	return &config, nil
}
