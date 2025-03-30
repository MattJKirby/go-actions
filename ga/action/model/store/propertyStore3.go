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
			WithUnsafeDecode[T](true),
		),
	}
}

func (aps *ActionPropertyStore[T]) NewProperty(property T) error {
	return aps.Insert(property.GetPropertyId(), &property)
}

func (aps *ActionPropertyStore[T]) MarshalJSON() ([]byte, error) {
	identifables := make([]*T, 0, len(aps.entries))
	for _, entry := range aps.entries {
		identifables = append(identifables, entry)
	}
	return json.Marshal(identifables)
}

func (aps *ActionPropertyStore[T]) UnmarshalJSON(data []byte) error {
	var identifables []T
	json.Unmarshal(data, &identifables)

	for _, item := range identifables {
		asdf, _ := json.Marshal(item)
		aps.unmarshalEntry(item.GetPropertyId(), asdf)
	}

	return nil
}
