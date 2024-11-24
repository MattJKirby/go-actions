package io

import (
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
