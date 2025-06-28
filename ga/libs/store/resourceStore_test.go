package store

import (
	"encoding/json"
	"go-actions/ga/utils/testing/assert"
	"testing"
)

type IdProp struct {
	Id    string
	Value string
}

func (p IdProp) getId() string {
	return p.Id
}

func TestNewResource(t *testing.T) {
	store := NewResourceStore(IdProp.getId, false)

	err := store.NewResource(IdProp{Id: "id", Value: "val"})

	assert.Equals(t, nil, err)
	assert.Equals(t, 1, len(store.Store.entries))
}

func TestGetDefault(t *testing.T) {
	existing := &IdProp{Id: "1", Value: "val"}
	defaultProp := &IdProp{Id: "2", Value: "val"}

	tests := []struct {
		name        string
		input       string
		expected    *IdProp
		expectedLen int
	}{
		{name: "existing Id - get path", input: "1", expected: existing, expectedLen: 1},
		{name: "non-existing Id - default path", input: "nonExisting", expected: defaultProp, expectedLen: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewResourceStore(IdProp.getId, false)
			store.NewResource(*existing)

			retrieved := store.GetDefault(test.input, func() IdProp {
				return *test.expected
			})

			assert.Equals(t, test.expected, &retrieved)
			assert.Equals(t, test.expectedLen, len(store.Store.entries))
		})
	}
}

func TestMarshalsResourceStore(t *testing.T) {
	store := NewResourceStore(IdProp.getId, false)
	store.NewResource(IdProp{Id: "id", Value: "val"})

	marshalled, err := json.Marshal(store)

	assert.Equals(t, nil, err)
	assert.Equals(t, `[{"Id":"id","Value":"val"}]`, string(marshalled))
}

func TestUnmarshalUpdate(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		unsafe     bool
		expected   IdProp
		expectedId string
		expectErr  bool
	}{
		{name: "valid - safe, existing resource", unsafe: false, input: `[{"Id":"id","Value":"1"}]`, expected: IdProp{"id", "1"}, expectedId: "id", expectErr: false},
		{name: "valid - unsafe, not existing resource", unsafe: true, input: `[{"Id":"x","Value":"1"}]`, expected: IdProp{"x", "1"}, expectedId: "x", expectErr: false},
		{name: "invalid - safe, not existing resource", unsafe: false, input: `[{"Id":"x","Value":"1"}]`, expected: IdProp{}, expectedId: "x", expectErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewResourceStore(IdProp.getId, test.unsafe)
			store.NewResource(IdProp{Id: "id", Value: "0"})

			err := store.UnmarshalJSON([]byte(test.input))
			assert.Equals(t, test.expectErr, err != nil)
			assert.Equals(t, test.expected, store.Store.entries[test.expectedId])
		})
	}
}
