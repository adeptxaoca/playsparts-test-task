package config

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"

	"part_handler/internal/pkg/validator"
)

type ServerConf struct {
	Port uint
}

type DatabaseConf struct {
	User string
	Pass string
	Addr string
	Name string
}

// App configuration structure
type Config struct {
	Server    ServerConf
	Database  DatabaseConf
	Json      jsoniter.API
	Validator *validator.Validator
}

// Basic configuration of the application and related components
func AppConfiguration(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	conf := Config{
		Server: ServerConf{Port: viper.GetUint("server.port")},
		Database: DatabaseConf{
			User: viper.GetString("database.user"),
			Pass: viper.GetString("database.pass"),
			Addr: viper.GetString("database.addr"),
			Name: viper.GetString("database.name"),
		},
		Json:      jsoniter.ConfigCompatibleWithStandardLibrary,
		Validator: validator.New(),
	}

	return &conf, nil
}
