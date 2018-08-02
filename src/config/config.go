package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

var config AppConfig

type AppConfig struct {
	Region          string `envconfig:"REGION"`
	AccessKeyID     string `envconfig:"ACCESS_KEY_ID"`
	SecretAccessKey string `envconfig:"SECRET_ACCESS_KEY"`
}

func init() {
	config = RefreshConfig()
}

func GetConfig() AppConfig {
	return config
}

func RefreshConfig() AppConfig {
	if err := envconfig.Process("app", &config); err != nil {
		panic(fmt.Sprintf("AppConfig file error: %s", err))
	}
	return config
}
