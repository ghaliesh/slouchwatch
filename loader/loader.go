package loader

type EnvFileLoader interface {
	Load(filenames ...string) error
}
