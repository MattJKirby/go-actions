package io

import (
	"encoding/json"
	"fmt"
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
	input := store.GetOrDefault("resource1", newInput)

	t.Run("test marshal", func(t *testing.T) {
		marshalled, _ := json.Marshal(store)
		marshalledinput, _ := json.Marshal(input)
		expected := fmt.Sprintf(`{"resource1":%s}`, marshalledinput)

		asserts.Equals(t, expected, string(marshalled))
	})
}
