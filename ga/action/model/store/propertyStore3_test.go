package store

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

type IdProp struct {
	Id string
	Value string
}

func (ip IdProp) GetPropertyId() string {
	return ip.Id
}

func TestNewProperty(t *testing.T) {
	store := NewActionPropertyStore[IdentifiableProperty]()

	err := store.NewProperty(&IdProp{Id: "id", Value: "val"})

	asserts.Equals(t, nil, err)
	asserts.Equals(t, 1, len(store.entries))
}

func TestMarshalx(t *testing.T){
	store := NewActionPropertyStore[IdentifiableProperty]()
	store.NewProperty(&IdProp{Id: "id", Value: "val"})

	marshalled , err := json.Marshal(store)

	asserts.Equals(t, nil, err)
	asserts.Equals(t, `[{"Id":"id","Value":"val"}]`, string(marshalled))
}

func TestUnmarshalx(t *testing.T){
	tests := []struct{
		name string
		input string
		expected *IdProp
		err bool
	}{
		{name: "valid, existing resource", input: `[{"Id":"id","Value":"1"}]`, expected: &IdProp{"id", "1"}, err: false},
		{name: "valid, not existing resource", input: `[{"Id":"x","Value":"1"}]`, expected: &IdProp{"x", "1"}, err: false},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewActionPropertyStore[IdProp]()
			store.NewProperty(IdProp{Id: "id", Value: "0"})

			err := store.UnmarshalJSON([]byte(test.input))
			asserts.Equals(t, test.err, err != nil)
			asserts.Equals(t, test.expected, store.entries[test.expected.Id])
		})
	}
}