package io

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestReferenceStoreAdd(t *testing.T) {
	store := NewActionReferenceStore[ActionInputReference]()

	t.Run("test store add", func(t *testing.T) {
		ref := NewActionInputReference("action", "input")
		store.Add(ref)

		asserts.Equals(t, 1, len(store.references))
	})
}