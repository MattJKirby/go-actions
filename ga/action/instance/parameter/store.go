package parameter

import (
	"encoding/json"
	"fmt"
)

type Store struct {
	parameters map[string]any
}

func NewStore() *Store {
	return &Store{
		parameters: make(map[string]any),
	}
}

func GetOrDefault[T any](name string, defaultValue T) func(*Store) *ActionParameter[T] {
	return func(s *Store) *ActionParameter[T] {
		_, exists := s.parameters[name]
		if !exists {
			s.parameters[name] = NewActionParameter(name, defaultValue)
		}

		return any(s.parameters[name]).(*ActionParameter[T])
	}
}

func (s *Store) MarshalJSON() ([]byte, error) {
	parameters := make(map[string]any)
	for name, value := range s.parameters {
		parameters[name] = value
	}

	return json.Marshal(parameters)
}

func (s *Store) UnmarshalJSON(data []byte) error {
	rawParameters := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &rawParameters)
	if err != nil {
		return err
	}

	for name, raw := range rawParameters {
		param, exists := s.parameters[name]
		if !exists {
			return fmt.Errorf("error unmashalling parameters: parameter '%s' does not exist", name)
		}

		err = json.Unmarshal(raw, param)
		if err != nil {
			return err
		}
	}
	return nil
}