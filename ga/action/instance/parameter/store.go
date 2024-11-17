package parameter

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type TypedParameter[T any] struct {
	parameterType reflect.Type
	parameterValue T
}

type Store struct {
	parameters map[string]*TypedParameter[any]
}

type MarshalledStore = map[string]any

func NewStore() *Store {
	return &Store{
		parameters: make(map[string]*TypedParameter[any]),
	}
}

func (s *Store) Get(name string) (*TypedParameter[any], error) {
	val, exists := s.parameters[name]
	if !exists {
		return nil, fmt.Errorf("no such parameter with name '%s'", name)
	}
	return val, nil
}

func GetOrDefault[T any](name string, defaultValue T) func(*Store) *ActionParameter[T] {
	return func(s *Store) *ActionParameter[T] {
		_, exists := s.parameters[name]
		if !exists {
			parameterValue := NewActionParameter(name, defaultValue)
			parameterType := reflect.TypeOf(defaultValue)
			s.parameters[name] = &TypedParameter[any]{parameterType, parameterValue}
		}

		return any(s.parameters[name].parameterValue).(*ActionParameter[T])
	}
}

func (s *Store) MarshalJSON() ([]byte, error) {
	parameters := make(map[string]any)
	for name,value := range s.parameters {
		parameters[name] = value.parameterValue
	}
	
	return json.Marshal(parameters)
}