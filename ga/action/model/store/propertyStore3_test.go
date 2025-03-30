package store

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

type IdProp struct {
	Id    string
	Value string
}

func (ip IdProp) GetPropertyId() string {
	return ip.Id
}

func TestNewProperty(t *testing.T) {
	store := NewActionPropertyStore[IdentifiableProperty](false)

	err := store.NewProperty(&IdProp{Id: "id", Value: "val"})

	asserts.Equals(t, nil, err)
	asserts.Equals(t, 1, len(store.entries))
}

func TestMarshalx(t *testing.T) {
	store := NewActionPropertyStore[IdentifiableProperty](false)
	store.NewProperty(&IdProp{Id: "id", Value: "val"})

	marshalled, err := json.Marshal(store)

	asserts.Equals(t, nil, err)
	asserts.Equals(t, `[{"Id":"id","Value":"val"}]`, string(marshalled))
}

func TestUnmarshalUpdate(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		unsafe     bool
		expected   *IdProp
		expectedId string
		expectErr  bool
	}{
		{name: "valid - safe, existing resource", unsafe: false, input: `[{"Id":"id","Value":"1"}]`, expected: &IdProp{"id", "1"}, expectedId: "id", expectErr: false},
		{name: "valid - unsafe, not existing resource", unsafe: true, input: `[{"Id":"x","Value":"1"}]`, expected: &IdProp{"x", "1"}, expectedId: "x", expectErr: false},
		{name: "invalid - safe, not existing resource", unsafe: false, input: `[{"Id":"x","Value":"1"}]`, expected: nil, expectedId: "x", expectErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewActionPropertyStore[IdProp](test.unsafe)
			store.NewProperty(IdProp{Id: "id", Value: "0"})

			err := store.UnmarshalJSON([]byte(test.input))
			asserts.Equals(t, test.expectErr, err != nil)
			asserts.Equals(t, test.expected, store.entries[test.expectedId])
		})
	}
}
