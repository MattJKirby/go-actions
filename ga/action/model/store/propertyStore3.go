package store

type IdentifiableProperty interface {
	GetPropertyId() string
}

type ActionPropertyStore[T IdentifiableProperty] struct {
	*BaseStore[T]
}

func NewActionPropertyStore[T IdentifiableProperty]() *ActionPropertyStore[T] {
	return &ActionPropertyStore[T]{
		BaseStore: NewBaseStore[T](),
	}
}

func (aps *ActionPropertyStore[T]) Store(property T) error {
	return aps.store(property.GetPropertyId(), &property)
}