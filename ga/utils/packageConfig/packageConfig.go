package packageConfig

type Option[T any] func(*T)

func NewPackageConfig[T any](cfg *T, opts ...Option[T]) *T {
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

