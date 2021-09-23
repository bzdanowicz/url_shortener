package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseAddress string
	DatabaseName    string
	CacheAddress    string
}

func LoadConfig(path string) (cfg *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return
}
