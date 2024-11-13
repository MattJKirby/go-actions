package parameter

import "encoding/json"

type Parameter interface {
	Value() any
	DefaultValue() any
	SetValue(value any)
}

type ActionParameter[T any] struct {
	name         string
	value        T
	defaultValue T
}

type MarshalledActionParameter[T any] struct {
	Name string
	Value T
}

func NewActionParameter[T any](Name string, DefaultValue T) *ActionParameter[T] {
	return &ActionParameter[T]{
		name:         Name,
		value:        DefaultValue,
		defaultValue: DefaultValue,
	}
}

func (ap *ActionParameter[T]) Value() T {
	return ap.value
}

func (ap *ActionParameter[T]) DefaultValue() T {
	return ap.defaultValue
}

func (ap *ActionParameter[T]) SetValue(value T) {
	ap.value = value
}

func (ap *ActionParameter[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&MarshalledActionParameter[T]{
		Name: ap.name,
		Value: ap.value,
	})
}
