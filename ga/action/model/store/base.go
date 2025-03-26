package store

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/utils/marshalling"
)

type BaseStore[T any] struct {
	entries map[string]*T
	config baseStoreConfig
}

type marshalledEntry[T any] struct {
	Id    string `json:"Id"`
	Value *T     `json:"Value"`
}

func NewBaseStore[T any](opts ...baseStoreOption[T]) *BaseStore[T] {
	store := &BaseStore[T]{
		entries: make(map[string]*T),
		config: baseStoreConfig{
			unsafeDecode: false,
		},
	}

	for _, opt := range opts {
		opt(store)
	}

	return store
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

func (bs *BaseStore[T]) MarshalJSON() ([]byte, error) {
	marshalledEntries := make([]*marshalledEntry[T], 0, len(bs.entries))
	for key, value := range bs.entries {
		marshalledEntries = append(marshalledEntries, &marshalledEntry[T]{key, value})
	}
	return json.Marshal(marshalledEntries)
}

func (bs *BaseStore[T]) UnmarshalJSON(data []byte) error {
	marshalledEntries := make([]*marshalledEntry[json.RawMessage], 0, len(bs.entries))
	if _, err := marshalling.StrictDecode(data, &marshalledEntries); err != nil {
		return err
	}

	for _, marshalledEntry := range marshalledEntries {
		if _, exists := bs.entries[marshalledEntry.Id]; !exists {
			return fmt.Errorf("failed to unmarshal: entry with identifier '%s' does not exist", marshalledEntry.Id)
		}

		if _, err := marshalling.StrictDecode(*marshalledEntry.Value, bs.entries[marshalledEntry.Id]); err != nil {
			return err
		}
	}
	return nil
}
