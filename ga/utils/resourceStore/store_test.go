package resourceStore

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestStoreGetOrDefault(t *testing.T) {
	type testStruct struct {
		val string
	}

	resource := &testStruct{val: "asdf"}

	t.Run("test get", func(t *testing.T) {
		store := NewStore[testStruct]()
		store.resources["name"] = resource
		got := store.GetOrDefault("name", resource)
		asserts.Equals(t, resource, got)
	})

	t.Run("test default", func(t *testing.T) {
		store := NewStore[testStruct]()
		got := store.GetOrDefault("not", resource)
		asserts.Equals(t, resource, got)
	})
}
