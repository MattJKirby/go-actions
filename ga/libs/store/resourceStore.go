package store

import "encoding/json"

type IdentifiableResource interface {
	GetResourceId() string
}

type ResourceStore[T IdentifiableResource] struct {
	Store *BaseStore[T]
}

func NewResourceStore[T IdentifiableResource](unsafeUpdate bool) *ResourceStore[T] {
	return &ResourceStore[T]{
		Store: NewBaseStore(
			WithUnsafeUpdate[T](unsafeUpdate),
		),
	}
}

func (aps *ResourceStore[T]) NewResource(property T) error {
	return aps.Store.Insert(property.GetResourceId(), &property)
}

func (aps *ResourceStore[T]) GetDefault(property T) T {
	return *aps.Store.GetDefault(property.GetResourceId(), func() *T { return &property })
}

func (aps *ResourceStore[T]) GetResource(propertyId string) (*T, error) {
	return aps.Store.Get(propertyId)
}

func (aps *ResourceStore[T]) MarshalJSON() ([]byte, error) {
	values := make([]*T, 0, len(aps.Store.entries))
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
		if err := aps.Store.Update(item.GetResourceId(), &item); err != nil {
			return err
		}
	}

	return nil
}
