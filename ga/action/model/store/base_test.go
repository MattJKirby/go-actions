package store

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

type prop struct {
	Name string
	Val string
}

func TestStore(t *testing.T){
	store := NewBaseStore[prop]()
	store.store("anon", &prop{})
	asserts.Equals(t, 1, len(store.entries))
}

func TestGet2(t *testing.T){
	existingProp := &prop{"id", "val"}
	tests := []struct{
		name string
		key string
		expect *prop
		err bool
	}{
		{name: "existing prop", key:"id", expect: existingProp, err: false},
		{name: "not existing prop", key: "badId", expect: nil, err: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewBaseStore[prop]()
			store.store(existingProp.Name, existingProp)

			retrieved, err := store.get(test.key)
			asserts.Equals(t, test.expect, retrieved)
			asserts.Equals(t, test.err, err != nil)
		})
	}
}

func TestGetDefault(t *testing.T){

  existing := &prop{"idA", "valA"}
  defaultProp := &prop{"idB", "valB"}
  
  tests := []struct{
    name string
    input string
    expected *prop
    expectedLen int
  }{
    {name: "existing Id - get path", input: "id", expected: existing, expectedLen: 1},
    {name: "non-existing Id - default path", input: "nonExisting", expected: defaultProp, expectedLen: 2},
  }

  for _,test := range tests {
    t.Run(test.name, func(t *testing.T) {
      store := NewBaseStore[prop]()
      store.store("id", existing)

      retrieved := store.getDefault(test.input, func() *prop {
        return defaultProp
      })

      asserts.Equals(t, test.expected, retrieved)
      asserts.Equals(t, test.expectedLen, len(store.entries))
    })
  }
}

