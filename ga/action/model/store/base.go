package store

import (
	"encoding/json"
	"fmt"
)

type BaseStore[T any] struct {
	entries map[string]*T
}

type marshalledEntry[T any] struct {
	Id string `json:"Id"`
	Value *T `json:"Value"`
}

func NewBaseStore[T any]() *BaseStore[T] {
	return &BaseStore[T]{
		entries: make(map[string]*T),
	}
}

func (bs *BaseStore[T]) store(key string, value *T) error {
	if _, exists := bs.entries[key]; exists {
		return fmt.Errorf("entry with key %s already exists", key)
	}
	bs.entries[key] = value
	return nil
}

func (bs *BaseStore[T]) get(key string) (*T, error) {
	if item, exists := bs.entries[key]; exists {
		return item, nil
	}
	return nil, fmt.Errorf("entry with key %s does not exist", key)
}

func (bs *BaseStore[T]) getDefault(key string, defaultFn func() *T) *T {
	if _, exists := bs.entries[key]; !exists {
		bs.entries[key] = defaultFn()
	}
	return bs.entries[key]
}

func (bs *BaseStore[T]) MarshalJSON() ([]byte, error){
	marshalledEntries := make([]*marshalledEntry[T], 0, len(bs.entries))
	for key, value := range bs.entries {
		marshalledEntries = append(marshalledEntries, &marshalledEntry[T]{key, value})
	}
	return json.Marshal(marshalledEntries)
}
