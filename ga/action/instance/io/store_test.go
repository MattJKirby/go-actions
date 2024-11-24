package io

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

type testResource struct {
}

func (tr testResource) Name() string {
	return ""
}

func (tr testResource) Id() string {
	return ""
}

func newTestResource(string, string) *testResource {
	return &testResource{}
}

func TestNewStore(t *testing.T) {
	store := NewStore[testResource]("uid")

	t.Run("test get", func(t *testing.T) {
		expected := newTestResource("name", "uid")
		input := store.GetOrDefaultResource("name", newTestResource)
		asserts.Equals(t, expected, input)
	})
}
