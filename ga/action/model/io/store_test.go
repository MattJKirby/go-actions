package io

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/cr"
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

func TestMarshalStore(t *testing.T) {
	store := NewStore[Input]()
	input := store.GetOrDefault("resource1", NewInput("name", "uid", false))
	marshalledinput, _ := json.Marshal(input)

	marshalled, err := json.Marshal(store)
	asserts.Equals(t, err, nil)
	asserts.Equals(t, fmt.Sprintf(`{"resource1":%s}`, marshalledinput), string(marshalled))
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
		store := NewStore[Input]()
		store.GetOrDefault("input", NewInput("input", "uid", false))
		err := json.Unmarshal([]byte(test.Input), store)

		hasErr := err != nil
		fmt.Println(err)
		asserts.Equals(t, test.Error, hasErr)
	})
}
