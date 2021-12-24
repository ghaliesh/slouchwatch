package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvLoader interface {
	Load(filenames ...string) error
}

type DotEnvLoader struct {
}

func (d *DotEnvLoader) Load(filenames ...string) error {
	return godotenv.Load(filenames...)
}

func NewEnvLoaderInstance() EnvLoader {
	return &DotEnvLoader{}
}

func GetEnvVar(loader EnvLoader, key string) (string, error) {
	err := loader.Load(".env")

	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}
