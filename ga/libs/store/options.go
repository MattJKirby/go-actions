package store

type baseStoreOption[T any] func(*BaseStore[T])

type baseStoreOptions struct {
	unsafeUpdate bool
}

func baseStoreDefaultOptions () baseStoreOptions {
	return baseStoreOptions{
		unsafeUpdate: false,
	}
}

func WithUnsafeUpdate[T any](enabled bool) func(*BaseStore[T]) {
	return func(bs *BaseStore[T]) {
		bs.options.unsafeUpdate = enabled
	}
}
