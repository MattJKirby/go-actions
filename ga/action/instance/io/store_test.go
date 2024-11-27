package io

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/cr"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewStore(t *testing.T) {
	store := NewStore[Input]("uid")

	t.Run("test default", func(t *testing.T) {
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

func TestUnmarshalStore(t *testing.T) {
	tests := []cr.TestCase[string, any]{
		{Name: "valid", Input: `{"input":{"name":"input","id":"id","output":{"actionUid":"","resourceName":""}}}`, Error: false},
		{Name: "valid no ref", Input: `{"input":{"name":"input","id":"id","output":null}}`, Error: false},
		{Name: "invalid key", Input: `{"inputx":{"name":"input","id":"id","output":{"actionUid":"","resourceName":""}}}`, Error: true},
		{Name: "invalid name", Input: `{"input":{"name":"inputx","id":"id","output":{"actionUid":"","resourceName":""}}}`, Error: true},
		{Name: "bad json", Input: `true`, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[string, any]) {
		store := NewStore[Input]("uid")
		store.GetOrDefault("input", newInput)
		err := json.Unmarshal([]byte(test.Input), store)

		hasErr := err != nil
		if test.Error != hasErr {
			t.Errorf("error unmarshalling store: got %v", err)
		}


	})
}
