package config

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

// App configuration structure
type Config struct {
	NetAddress string
	ConnString string
	Json       jsoniter.API
}

// Basic configuration of the application and related components
func AppConfiguration(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	conf := Config{
		NetAddress: fmt.Sprintf("localhost:%s", viper.GetString("server.port")),
		ConnString: fmt.Sprintf("postgres://%s:%s@%s/%s",
			viper.GetString("database.user"),
			viper.GetString("database.pass"),
			viper.GetString("database.addr"),
			viper.GetString("database.name")),
		Json: jsoniter.ConfigCompatibleWithStandardLibrary,
	}

	return &conf, nil
}
