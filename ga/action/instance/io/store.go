package io

import (
	"encoding/json"
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

// func (s *Store[T]) UnmarshalJSON(data []byte) error {
// 	rawInputs := make(map[string]*T)
// 	err := json.Unmarshal(data, &rawInputs)
// 	if err != nil {
// 		return err
// 	}

// 	// for name, raw := range rawInputs {
// 	// 	_, exists := s.resources[name]
// 	// 	if !exists {
// 	// 		return fmt.Errorf("error unmashalling parameters: parameter '%s' does not exist", name)
// 	// 	}

// 	// 	s.resources[name] = raw
// 	// }
// 	fmt.Println("AAAAA")
// 	s.resources = rawInputs

// 	return nil
// }
