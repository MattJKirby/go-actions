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
