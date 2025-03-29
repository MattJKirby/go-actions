package store

type IdentifiableProperty interface {
	GetPropertyId() string
}

type ActionPropertyStore[T IdentifiableProperty] struct {
	*BaseStore[T]
}

func NewActionPropertyStore[T IdentifiableProperty]() *ActionPropertyStore[T] {
	return &ActionPropertyStore[T]{
		NewBaseStore[T](),
	}
}

func (aps *ActionPropertyStore[T]) NewProperty(property T) error {
	return aps.insert(property.GetPropertyId(), &property)
}

func (aps *ActionPropertyStore[T]) GetOrDefaultProperty(id string, defaultFn func() *T) *T {
	return aps.getDefault(id, defaultFn)
}

