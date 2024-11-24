package io

import "go-actions/ga/utils/resourceStore"

type ioResource interface {
	Name () string
	Id () string
}


type Store[T ioResource] struct {
	actionUid string
	resources *resourceStore.Store[T]
}

func NewStore[T ioResource](actionUid string) *Store[T] {
	return &Store[T]{
		actionUid: actionUid,
		resources: resourceStore.NewStore[T](),
	}
}
func (s *Store[T]) GetOrDefault(name string, ctor func(string, string) *T) *T {
		resource := ctor(name, s.actionUid)
		return s.resources.GetOrDefault(name, resource)
}