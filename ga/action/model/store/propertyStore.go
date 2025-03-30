package store

import "encoding/json"

type IdentifiableProperty interface {
	GetPropertyId() string
}

type PropertyStore[T IdentifiableProperty] struct {
	*BaseStore[T]
}

func NewPropertyStore[T IdentifiableProperty](unsafeUpdate bool) *PropertyStore[T] {
	return &PropertyStore[T]{
		NewBaseStore(
			WithUnsafeUpdate[T](unsafeUpdate),
		),
	}
}

func (aps *PropertyStore[T]) NewProperty(property T) error {
	return aps.Insert(property.GetPropertyId(), &property)
}

func (aps *PropertyStore[T]) MarshalJSON() ([]byte, error) {
	values := make([]*T, 0, len(aps.entries))
	for _, entry := range aps.entries {
		values = append(values, entry)
	}
	return json.Marshal(values)
}

func (aps *PropertyStore[T]) UnmarshalJSON(data []byte) error {
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
