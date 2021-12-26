package config

type EnvConfig interface {
	GetEnvVar(key string) (string, error)
}
