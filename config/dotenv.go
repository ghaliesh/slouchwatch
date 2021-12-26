package config

import (
	"os"

	"github.com/ghaliesh/slouchwatch/loader"
)

var envLoader loader.EnvFileLoader

type dotEnvConfig struct {
}

func (d *dotEnvConfig) GetEnvVar(key string) (string, error) {
	err := envLoader.Load(".env")

	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}

func NewDotEnvConfigInstance(fileLoader loader.EnvFileLoader) EnvConfig {
	envLoader = fileLoader
	return &dotEnvConfig{}
}
