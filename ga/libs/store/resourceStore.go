package store

import "encoding/json"

type IdentifiableResource interface {
	GetResourceId() string
}

type ResourceStore[T IdentifiableResource] struct {
	*BaseStore[T]
}

func NewResourceStore[T IdentifiableResource](unsafeUpdate bool) *ResourceStore[T] {
	return &ResourceStore[T]{
		NewBaseStore(
			WithUnsafeUpdate[T](unsafeUpdate),
		),
	}
}

func (aps *ResourceStore[T]) NewResource(property T) error {
	return aps.Insert(property.GetResourceId(), &property)
}

func (aps *ResourceStore[T]) MarshalJSON() ([]byte, error) {
	values := make([]*T, 0, len(aps.entries))
	for _, entry := range aps.entries {
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
		if err := aps.Update(item.GetResourceId(), &item); err != nil {
			return err
		}
	}

	return nil
}
