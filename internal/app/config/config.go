package config

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"

	"part_handler/internal/pkg/validator"
)

type DatabaseConf struct {
	User string
	Pass string
	Addr string
	Name string
	Url  string

	MaxConns int32
}

// App configuration structure
type Config struct {
	Database  DatabaseConf
	Json      jsoniter.API
	Validator *validator.Validator
}

// Basic configuration of the application and related components
func AppConfiguration() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := Config{
		Database: DatabaseConf{
			User:     viper.GetString("DATABASE_USER"),
			Pass:     viper.GetString("DATABASE_PASS"),
			Addr:     viper.GetString("DATABASE_ADDR"),
			Name:     viper.GetString("DATABASE_NAME"),
			Url:      viper.GetString("DATABASE_URL"),
			MaxConns: viper.GetInt32("DATABASE_MAX_CONNS"),
		},
		Json:      jsoniter.ConfigCompatibleWithStandardLibrary,
		Validator: validator.New(),
	}

	return &conf, nil
}
