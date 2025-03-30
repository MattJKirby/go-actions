package store

import "encoding/json"

type IdentifiableProperty interface {
	GetPropertyId() string
}

type ActionPropertyStore[T IdentifiableProperty] struct {
	*BaseStore[T]
}

func NewActionPropertyStore[T IdentifiableProperty]() *ActionPropertyStore[T] {
	return &ActionPropertyStore[T]{
		NewBaseStore(
			WithUnsafeUpdate[T](true),
		),
	}
}

func (aps *ActionPropertyStore[T]) NewProperty(property T) error {
	return aps.Insert(property.GetPropertyId(), &property)
}

func (aps *ActionPropertyStore[T]) MarshalJSON() ([]byte, error) {
	values := make([]*T, 0, len(aps.entries))
	for _, entry := range aps.entries {
		values = append(values, entry)
	}
	return json.Marshal(values)
}

func (aps *ActionPropertyStore[T]) UnmarshalJSON(data []byte) error {
	var values []T
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}

	for _, item := range values {
		if err := aps.Update(item.GetPropertyId(), &item); err != nil {
			return err
		}
	}

	return nil
}
