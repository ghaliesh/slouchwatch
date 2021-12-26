package loader

import (
	"github.com/joho/godotenv"
)

type dotEnvFileLoader struct {
}

func (d *dotEnvFileLoader) Load(filenames ...string) error {
	return godotenv.Load(filenames...)
}

func NewDotEnvFileLoader() EnvFileLoader {
	return &dotEnvFileLoader{}
}
