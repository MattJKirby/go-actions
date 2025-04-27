package store

import (
	"encoding/json"
	"go-actions/ga/utils/testing/assert"
	"testing"
)

type prop struct {
	Val string
}

func TestInsert(t *testing.T) {
	existingProp := &prop{"val"}
	tests := []struct {
		name        string
		input       string
		expectedLen int
		err         bool
	}{
		{name: "valid - non-existing key", input: "non existing", expectedLen: 2, err: false},
		{name: "invalid - key exists", input: "exists", expectedLen: 1, err: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewBaseStore[prop]()
			store.Insert("exists", existingProp)

			err := store.Insert(test.input, &prop{})
			assert.Equals(t, test.expectedLen, len(store.entries))
			assert.Equals(t, test.err, err != nil)
		})
	}
}

func TestGet2(t *testing.T) {
	existingProp := &prop{"value"}
	tests := []struct {
		name   string
		key    string
		expect *prop
		err    bool
	}{
		{name: "existing prop", key: "id", expect: existingProp, err: false},
		{name: "not existing prop", key: "badId", expect: nil, err: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewBaseStore[prop]()
			store.Insert("id", existingProp)

			retrieved, err := store.Get(test.key)
			assert.Equals(t, test.expect, retrieved)
			assert.Equals(t, test.err, err != nil)
		})
	}
}

func TestGetDefault(t *testing.T) {
	existing := &prop{"idA"}
	defaultProp := &prop{"idB"}

	tests := []struct {
		name        string
		input       string
		expected    *prop
		expectedLen int
	}{
		{name: "existing Id - get path", input: "id", expected: existing, expectedLen: 1},
		{name: "non-existing Id - default path", input: "nonExisting", expected: defaultProp, expectedLen: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewBaseStore[prop]()
			store.Insert("id", existing)

			retrieved := store.GetDefault(test.input, func() *prop {
				return defaultProp
			})

			assert.Equals(t, test.expected, retrieved)
			assert.Equals(t, test.expectedLen, len(store.entries))
		})
	}
}

func TestMarshal(t *testing.T) {
	store := NewBaseStore[prop]()
	store.Insert("id", &prop{"val"})

	marshalled, err := json.Marshal(store)

	assert.Equals(t, []byte(`[{"Id":"id","Value":{"Val":"val"}}]`), marshalled)
	assert.Equals(t, nil, err)
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValue *prop
		err           bool
	}{
		{name: "valid, entry already exists", input: `[{"Id":"id","Value":{"Val":"1"}}]`, expectedValue: &prop{"1"}, err: false},
		{name: "invalid, entry doesnt exist", input: `[{"Id":"x","Value":{"Val":"1"}}]`, expectedValue: &prop{"0"}, err: true},
		{name: "invalid, wrong entry json", input: `[{"X":"x","Value":{"Val":"1"}}]`, expectedValue: &prop{"0"}, err: true},
		{name: "invalid, wrong value json", input: `[{"Id":"id","Value":{"X":"1"}}]`, expectedValue: &prop{"0"}, err: true},
		{name: "invalid, bad store json", input: `0`, expectedValue: &prop{"0"}, err: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewBaseStore[prop]()
			store.Insert("id", &prop{"0"})

			err := json.Unmarshal([]byte(test.input), store)

			assert.Equals(t, test.err, err != nil)
			assert.Equals(t, test.expectedValue, store.entries["id"])
		})
	}
}

func TestUnsafeDecodeMarshal(t *testing.T) {
	store := NewBaseStore(
		WithUnsafeUpdate[prop](true),
	)

	err := json.Unmarshal([]byte(`[{"id":"id","Value":{"Val":"val"}}]`), store)
	assert.Equals(t, nil, err)
	assert.Equals(t, &prop{"val"}, store.entries["id"])
}
