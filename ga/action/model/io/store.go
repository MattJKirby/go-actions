package io

import (
	"encoding/json"
	"go-actions/ga/action/model/resources"
)

type Store[T any] struct {
	actionUid string
	resources *resources.ResourceStore[T]
}

func NewStore[T any](actionUid string) *Store[T] {
	return &Store[T]{
		actionUid: actionUid,
		resources: resources.NewResourceStore[T](),
	}
}

func (s *Store[T]) GetOrDefault(name string, ctor func(string, string) *T) *T {
	defaultVal := ctor(name, s.actionUid)
	return s.resources.GetOrDefault(name, defaultVal)
}

func (s *Store[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.resources)
}

func (s *Store[T]) UnmarshalJSON(data []byte) error {
	return s.resources.UnmarshalJSON(data)
}
