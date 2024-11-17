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
	for name,value := range s.parameters {
		parameters[name] = value
	}
	
	return json.Marshal(parameters)
}

func (s *Store) UnmarshalJSON(data []byte) error {
	parameters := make(map[string]any)
	err := json.Unmarshal(data, &parameters)
	if err != nil { 
		return err
	}

	for name,val := range parameters {
		param, exists := s.parameters[name]
		if !exists {
			return fmt.Errorf("error unmashalling parameters: parameter '%s' does not exist", name)
		}

		rawParam, err := json.Marshal(val)
		if err != nil {
			return err
		}

		err = json.Unmarshal(rawParam, param)
		if err != nil {
			return err
		}
	}
	return nil
}