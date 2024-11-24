package io

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewStore(t *testing.T) {
	store := NewStore[Input]("uid")

	t.Run("test get", func(t *testing.T) {
		expected := newInput("name", "uid")
		input := store.GetOrDefault("name", newInput)
		asserts.Equals(t, expected, input)
	})
}

func TestMarshalStore(t *testing.T) {
	store := NewStore[Input]("uid")
	store.GetOrDefault("resource1", newInput)

	t.Run("test marshal", func(t *testing.T) {
		marshalled, _ := json.Marshal(store)
		asserts.Equals(t, `{"resource1":{"name":"resource1","id":"uid__Input:resource1"}}`, string(marshalled))
	})
}
