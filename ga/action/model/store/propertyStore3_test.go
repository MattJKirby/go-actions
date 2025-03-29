package store

import (
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

func TestInsertion(t *testing.T) {
	store := NewActionPropertyStore[IdentifiableProperty]()

	err := store.Store(&IdProp{Id: "id", Value: "val"})

	asserts.Equals(t, nil, err)
	asserts.Equals(t, 1, len(store.entries))
}