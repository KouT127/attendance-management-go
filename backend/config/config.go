package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Database struct {
		Dbms   string
		User   string
		Pass   string
		Host   string
		Port   string
		DbName string
		Option string
	}

	Application struct {
		IsDebug bool
	}
}

type environment int

const (
	Development environment = iota
	Staging
	Production
)

func (e environment) path() string {
	switch e {
	case Staging:
		return "./backend/config/staging"
	case Production:
		return "./backend/config/production"
	default:
		return "./backend/config/development"
	}
}

var c Config

func Init(e environment) {
	viper.AddConfigPath(e.path())
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file read error")
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&c); err != nil {
		fmt.Println("config file Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewConfig() *Config {
	return &c
}
