// Package config provides all mapped env variables into struct
package config

import (
	"log/slog"
	"sync"

	"github.com/spf13/viper"
)

// Config schema for mapping env variables
type Config struct {
	Port string `mapstructure:"PORT"`

	DbHost  string `mapstructure:"DB_HOST"`
	DbPort  string `mapstructure:"DB_PORT"`
	DbUser  string `mapstructure:"DB_USER"`
	DbPass  string `mapstructure:"DB_PASS"`
	DbName  string `mapstructure:"DB_NAME"`
	DbDebug bool   `mapstructure:"DB_DEBUG"`
}

var (
	cfg  *Config
	lock = &sync.Mutex{}
)

func loadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var conf Config
	err = viper.Unmarshal(&conf)

	return &conf, err
}

// GetConfig returns config data
func GetConfig(path string) *Config {
	lock.Lock()
	defer lock.Unlock()

	if cfg == nil {
		var err error

		cfg, err = loadConfig(path)
		if err != nil {
			slog.Error("error LoadConfig", "err", err)
			panic(err)
		}
	}

	return cfg
}
