package parameter

import (
	"encoding/json"
	"fmt"
)

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

type marshalledActionParameter[T any] struct {
	Name  string `json:"name"`
	Value T `json:"value"`
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
	return json.Marshal(&marshalledActionParameter[T]{
		Name:  ap.name,
		Value: ap.value,
	})
}

func (ap *ActionParameter[T]) UnmarshalJSON(data []byte) error {
	var marshalled marshalledActionParameter[T]
	if err := json.Unmarshal(data, &marshalled); err != nil {
		return err
	}

	if marshalled.Name != ap.name {
		return fmt.Errorf("failed to unmarshal action parameter: '%s': name '%s' does not match expected '%s'", ap.name, marshalled.Name, ap.name)
	}

	ap.SetValue(marshalled.Value)
	return nil
}
