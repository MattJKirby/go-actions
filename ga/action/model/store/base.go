package store

import "fmt"

type BaseStore[T any] struct {
	entries map[string]*T
}

func NewBaseStore[T any]() *BaseStore[T]{
	return &BaseStore[T]{
		entries: make(map[string]*T),
	}
}

func (bs *BaseStore[T]) store(key string, value *T){
	bs.entries[key] = value
}

func (bs *BaseStore[T]) get(key string) (*T, error){
	if item, exists := bs.entries[key]; exists {
		return item, nil
	}
	return nil, fmt.Errorf("entry with key %s does not exist", key)
}

func (bs *BaseStore[T]) getDefault(key string, defaultFn func() *T) *T{
	if _, exists := bs.entries[key]; !exists {
		bs.entries[key] = defaultFn()
	}
	return bs.entries[key]
}