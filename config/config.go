package config

import (
	"log/slog"
	"sync"

	"github.com/spf13/viper"
)

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

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}

func GetConfig() *Config {
	lock.Lock()
	defer lock.Unlock()

	if cfg == nil {
		err := LoadConfig("../")
		if err != nil {
			slog.Error("error LoadConfig", "err", err)
			return nil
		}
	}

	return cfg
}
