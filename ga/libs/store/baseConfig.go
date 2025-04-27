package store

type baseStoreOption[T any] func(*BaseStore[T])

type baseStoreConfig struct {
	unsafeUpdate bool
}

func WithUnsafeUpdate[T any](enabled bool) func(*BaseStore[T]) {
	return func(bs *BaseStore[T]) {
		bs.config.unsafeUpdate = enabled
	}
}
