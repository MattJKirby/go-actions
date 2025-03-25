package store

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

type anonProp struct {
	Name string
	Val string
}

func TestStore(t *testing.T){
	store := NewBaseStore[anonProp]()
	store.store("anon", &anonProp{})
	asserts.Equals(t, 1, len(store.entries))
}

func TestGet2(t *testing.T){
	existingProp := &anonProp{"id", "val"}
	tests := []struct{
		name string
		key string
		expect *anonProp
		err bool
	}{
		{name: "existing prop", key:"id", expect: existingProp, err: false},
		{name: "not existing prop", key: "badId", expect: nil, err: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewBaseStore[anonProp]()
			store.store(existingProp.Name, existingProp)

			retrieved, err := store.get(test.key)
			asserts.Equals(t, test.expect, retrieved)
			asserts.Equals(t, test.err, err != nil)
		})
	}
}

