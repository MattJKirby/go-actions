package io

import (
	"go-actions/ga/action/model/resources"
)

type IOStore[T any] struct {
	actionUid string
	*resources.ResourceStore[T]
}

func NewIOStore[T any](actionUid string) *IOStore[T] {
	return &IOStore[T]{
		actionUid,
		resources.NewResourceStore[T](),
	}
}

func (i *IOStore[T]) GetOrDefault(name string, ctor func(string, string) *T) *T {
	defaultVal := ctor(name, i.actionUid)
	return i.ResourceStore.GetOrDefault(name, defaultVal)
}
