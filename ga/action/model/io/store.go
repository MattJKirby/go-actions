package io

import (
	"go-actions/ga/action/model/resources"
)

type Store[T any] struct {
	actionUid string
	*resources.ResourceStore[T]
}

func NewStore[T any](actionUid string) *Store[T] {
	return &Store[T]{
		actionUid,
		resources.NewResourceStore[T](),
	}
}
