package io

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/cr"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestStoreGetOrDefault(t *testing.T) {
	store := NewIOStore[Input]("uid")
	expectedInput := newInput("name", "uid")

	t.Run("test default", func(t *testing.T) {
		input := store.GetOrDefault("name", newInput)
		asserts.Equals(t, expectedInput, input)
	})
}

func TestMarshalStore(t *testing.T) {
	store := NewIOStore[Input]("uid")
	input := store.GetOrDefault("resource1", newInput)
	marshalledinput, _ := json.Marshal(input)

	t.Run("test marshal", func(t *testing.T) {
		marshalled, err := json.Marshal(store)
		asserts.Equals(t, err, nil)
		asserts.Equals(t, fmt.Sprintf(`{"resource1":%s}`, marshalledinput), string(marshalled))
	})
}

func TestUnmarshalStore(t *testing.T) {
	tests := []cr.TestCase[string, int]{
		{Name: "valid", Input: `{"input":{"name":"input","id":"id","output":{"actionUid":"","resourceName":""}}}`, Error: false},
		{Name: "valid no ref", Input: `{"input":{"name":"input","id":"id","output":null}}`, Error: false},
		{Name: "invalid key", Input: `{"inputx":{"name":"input","id":"id","output":{"actionUid":"","resourceName":""}}}`, Error: true},
		{Name: "invalid name", Input: `{"input":{"name":"inputx","id":"id","output":{"actionUid":"","resourceName":""}}}`, Error: true},
		{Name: "bad store json ", Input: `true`, Error: true},
		{Name: "bad resource json", Input: `{"input":{"name":true,"id":"id","output":{"actionUid":"","resourceName":""}}}`, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[string, int]) {
		store := NewIOStore[Input]("uid")
		store.GetOrDefault("input", newInput)
		err := json.Unmarshal([]byte(test.Input), store)

		hasErr := err != nil
		fmt.Println(err)
		asserts.Equals(t, test.Error, hasErr)
	})
}
