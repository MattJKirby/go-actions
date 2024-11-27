package io

import (
	"encoding/json"
	"fmt"
)

type Store[T any] struct {
	actionUid string
	resources map[string]*T
}

func NewStore[T any](actionUid string) *Store[T] {
	return &Store[T]{
		actionUid: actionUid,
		resources: make(map[string]*T),
	}
}

func (s *Store[T]) GetOrDefault(name string, ctor func(string, string) *T) *T {
	_, exists := s.resources[name]
	if !exists {
		s.resources[name] = ctor(name, s.actionUid)
	}

	return s.resources[name]
}

func (s *Store[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.resources)
}

func (s *Store[T]) UnmarshalJSON(data []byte) error {
	rawInputs := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &rawInputs); err != nil {
		return err
	}

	for name, raw := range rawInputs {
		if _, exists := s.resources[name]; !exists {
			return fmt.Errorf("error unmashalling resource: resource '%s' does not exist", name)
		}

		if err := json.Unmarshal(raw, s.resources[name]); err != nil {
			return err
		}
	}

	return nil
}
