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
	existingProp := prop{"val"}
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

			err := store.Insert(test.input, prop{})
			assert.Equals(t, test.expectedLen, len(store.entries))
			assert.Equals(t, test.err, err != nil)
		})
	}
}

func TestGet2(t *testing.T) {
	existingProp := prop{"value"}
	tests := []struct {
		name   string
		key    string
		expect prop
		err    bool
	}{
		{name: "existing prop", key: "id", expect: existingProp, err: false},
		{name: "not existing prop", key: "badId", expect: prop{}, err: true},
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

func TestGetEntries(t *testing.T) {
	store := NewBaseStore[prop]()
	p1 := prop{"1"}
	p2 := prop{"2"}
	store.Insert("1", p1)
	store.Insert("2", p2)

	expected := map[string]prop{"1": p1, "2": p2}

	assert.Equals(t, expected, store.GetEntries())
}

func TestMarshal(t *testing.T) {
	store := NewBaseStore[prop]()
	store.Insert("id", prop{"val"})

	marshalled, err := json.Marshal(store)

	assert.Equals(t, []byte(`[{"Id":"id","Value":{"Val":"val"}}]`), marshalled)
	assert.Equals(t, nil, err)
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValue prop
		err           bool
	}{
		{name: "valid, entry already exists", input: `[{"Id":"id","Value":{"Val":"1"}}]`, expectedValue: prop{"1"}, err: false},
		{name: "invalid, entry doesnt exist", input: `[{"Id":"x","Value":{"Val":"1"}}]`, expectedValue: prop{"0"}, err: true},
		{name: "invalid, wrong entry json", input: `[{"X":"x","Value":{"Val":"1"}}]`, expectedValue: prop{"0"}, err: true},
		{name: "invalid, wrong value json", input: `[{"Id":"id","Value":{"X":"1"}}]`, expectedValue: prop{"0"}, err: true},
		{name: "invalid, bad store json", input: `0`, expectedValue: prop{"0"}, err: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewBaseStore[prop]()
			store.Insert("id", prop{"0"})

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
	assert.Equals(t, prop{"val"}, store.entries["id"])
}
