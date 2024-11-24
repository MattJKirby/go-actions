package reference

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestReferenceStoreAdd(t *testing.T) {
	store := NewActionReferenceStore[Input]()

	t.Run("test store add", func(t *testing.T) {
		ref := NewInput("action", "input")
		store.Add(ref)

		asserts.Equals(t, 1, len(store.references))
	})
}

func TestStoreGetOrDefault(t *testing.T) {
	ref := NewInput("ref", "input")

	t.Run("test get", func(t *testing.T) {
		store := NewActionReferenceStore[Input]()
		store.Add(ref)
		got := store.GetOrDefault(ref)
		asserts.Equals(t, ref, got)
	})

	t.Run("test default", func(t *testing.T) {
		store := NewActionReferenceStore[Input]()
		got := store.GetOrDefault(ref)
		asserts.Equals(t, ref, got)
	})
}
