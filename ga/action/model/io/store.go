package io

import (
	"go-actions/ga/action/model/resources"
)

type Store[T any] struct {
	*resources.ResourceStore[T]
}

func NewStore[T any]() *Store[T] {
	return &Store[T]{
		resources.NewResourceStore[T](),
	}
}
