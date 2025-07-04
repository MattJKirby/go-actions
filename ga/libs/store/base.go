package store

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/utils/marshalling"
)

type BaseStore[T any] struct {
	entries map[string]T
	options baseStoreOptions
}

type marshalledEntry[T any] struct {
	Id    string `json:"Id"`
	Value T     `json:"Value"`
}

func NewBaseStore[T any](opts ...baseStoreOption[T]) *BaseStore[T] {
	store := &BaseStore[T]{
		entries: make(map[string]T),
		options: baseStoreDefaultOptions(),
	}

	for _, opt := range opts {
		opt(store)
	}

	return store
}

func (bs *BaseStore[T]) Insert(key string, value T) error {
	if _, exists := bs.entries[key]; exists {
		return fmt.Errorf("entry with key %s already exists", key)
	}
	bs.entries[key] = value
	return nil
}

func (bs *BaseStore[T]) Get(key string) (T, error) {
	if item, exists := bs.entries[key]; exists {
		return item, nil
	}
	return *new(T), fmt.Errorf("entry with key %s does not exist", key)
}

func (bs *BaseStore[T]) Update(key string, value T) error {
	_, exists := bs.entries[key]
	if !exists && !bs.options.unsafeUpdate {
		return fmt.Errorf("failed to unmarshal: entry with identifier '%s' does not exist", key)
	}

	bs.entries[key] = value
	return nil
}

func (bs *BaseStore[T]) GetEntries() map[string]T {
	entries := make(map[string]T)
	for name, item := range bs.entries {
		entries[name] = item
	}
	return entries
}

func (bs *BaseStore[T]) MarshalJSON() ([]byte, error) {
	marshalledEntries := make([]*marshalledEntry[T], 0, len(bs.entries))
	for key, value := range bs.entries {
		marshalledEntries = append(marshalledEntries, &marshalledEntry[T]{key, value})
	}
	return json.Marshal(marshalledEntries)
}

func (bs *BaseStore[T]) UnmarshalJSON(data []byte) error {
	var marshalledEntries []*marshalledEntry[T]
	if _, err := marshalling.StrictDecode(data, &marshalledEntries); err != nil {
		return err
	}

	for _, entry := range marshalledEntries {
		if err := bs.Update(entry.Id, entry.Value); err != nil {
			return err
		}
	}
	return nil
}
