package io

import (
	"encoding/json"
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
		input := store.GetOrDefault("name", newTestResource)
		asserts.Equals(t, expected, input)
	})
}

func TestMarshalStore(t *testing.T) {
	store := NewStore[testResource]("uid")
	store.GetOrDefault("resource1", newTestResource)

	t.Run("test marshal", func(t *testing.T) {
		marshalled, _ := json.Marshal(store)
		asserts.Equals(t, `{"resource1":{}}`, string(marshalled))
	})
}
