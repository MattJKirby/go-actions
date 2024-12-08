package io

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestStoreGetOrDefault(t *testing.T) {
	store := NewStore[Input]()
	expectedInput := NewInput("name", "uid", false)

	t.Run("test default", func(t *testing.T) {
		input := store.GetOrDefault("name", expectedInput)
		asserts.Equals(t, expectedInput, input)
	})
}
