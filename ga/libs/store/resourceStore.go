package store

import "encoding/json"

type Identifiable interface {
	GetId() string
}

type ResourceStore[T any] struct {
	Store         *BaseStore[T]
	GetResourceId func(T) string
}

func NewResourceStore[T any](id func(T) string, unsafeUpdate bool) *ResourceStore[T] {
	return &ResourceStore[T]{
		Store: NewBaseStore(
			WithUnsafeUpdate[T](unsafeUpdate),
		),
		GetResourceId: id,
	}
}

func (aps *ResourceStore[T]) NewResource(property T) error {
	return aps.Store.Insert(aps.GetResourceId(property), property)
}

func (aps *ResourceStore[T]) GetDefault(key string, defaultFn func() T) T {
	if exists, err := aps.Store.Get(key); err == nil {
		return exists
	}

	defaultResource := defaultFn()
	aps.NewResource(defaultResource)
	return defaultResource
}

func (aps *ResourceStore[T]) GetResource(propertyId string) (T, error) {
	return aps.Store.Get(propertyId)
}

func (aps *ResourceStore[T]) MarshalJSON() ([]byte, error) {
	values := make([]T, 0, len(aps.Store.entries))
	for _, entry := range aps.Store.entries {
		values = append(values, entry)
	}
	return json.Marshal(values)
}

func (aps *ResourceStore[T]) UnmarshalJSON(data []byte) error {
	var values []T
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}

	for _, item := range values {
		if err := aps.Store.Update(aps.GetResourceId(item), item); err != nil {
			return err
		}
	}

	return nil
}
