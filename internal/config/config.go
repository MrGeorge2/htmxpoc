package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var conf *Config

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	HostName string
	Port     int
	UserName string
	Password string
	DbName   string
}

func (d DatabaseConfig) CreateConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", d.HostName, d.Port, d.UserName, d.Password)
}

func InitConfig() error {
	if conf != nil {
		return nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	conf = &Config{}

	return viper.Unmarshal(conf)
}

func GetConfig() (Config, error) {
	if conf == nil {
		InitConfig()
	}

	return *conf, nil
}
