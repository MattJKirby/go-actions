package references

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

func TestStoreGetOrDefault(t *testing.T) {
	ref := NewActionInputReference("ref", "input")

	t.Run("test get", func(t *testing.T) {
		store := NewActionReferenceStore[ActionInputReference]()
		store.Add(ref)
		got := store.GetOrDefault(ref)
		asserts.Equals(t, ref, got)
	})

	t.Run("test default", func(t *testing.T) {
		store := NewActionReferenceStore[ActionInputReference]()
		got := store.GetOrDefault(ref)
		asserts.Equals(t, ref, got)
	})
}