package store

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/cr"
	"go-actions/ga/cr/asserts"
	"testing"
)

type testProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func TestAdd(t *testing.T) {
	store := NewPropertyStore[testProperty]()

	store.Add("r", &testProperty{})
	asserts.Equals(t, 1, len(store.properties))
}

func TestGet(t *testing.T) {
	store := NewPropertyStore[testProperty]()
	prop := &testProperty{}
	store.Add("r", prop)

	t.Run("test valid", func(t *testing.T) {
		val, err := store.Get("r")
		asserts.Equals(t, err, nil)
		asserts.Equals(t, prop, val)
	})

	t.Run("test invalid", func(t *testing.T) {
		val, err := store.Get("x")
		asserts.Equals(t, true, err != nil)
		asserts.Equals(t, nil, val)
	})
}

func TestGetOrDefault(t *testing.T) {
	prop := &testProperty{"r", "v"}

	t.Run("test get path", func(t *testing.T) {
		store := NewPropertyStore[testProperty]()
		store.Add(prop.Name, prop)
		retrieved := store.GetOrDefault(prop.Name, func() *testProperty {
			return &testProperty{"r", "v"}
		})
		asserts.Equals(t, prop, retrieved)
	})

	t.Run("test default path", func(t *testing.T) {
		store := NewPropertyStore[testProperty]()
		retrieved := store.GetOrDefault("someProp", func() *testProperty {
			return &testProperty{"r", "v"}
		})
		asserts.Equals(t, prop, retrieved)
	})
}

func TestCustomMarshalling(t *testing.T) {
	store := NewPropertyStore[testProperty]()
	prop := &testProperty{"r", "v"}
	marshalledProp, _ := json.Marshal(prop)
	store.Add("r", prop)

	expected := fmt.Sprintf(`[{"propertyId":"r","property":%s}]`, string(marshalledProp))
	marshalledStore, err := json.Marshal(store)
	asserts.Equals(t, nil, err)
	asserts.Equals(t, []byte(expected), marshalledStore)
}

func TestCustomUnmarshalling(t *testing.T) {

	existingProp := &testProperty{"r", "v"}

	tests := []cr.TestCase[string, *testProperty]{
		{Name: "valid json, existing prop", Input: `[{"propertyId":"r","property":{"name":"a","value":"b"}}]`, Expected: &testProperty{"a", "b"}, Error: false},
		{Name: "valid json, non-existing prop id", Input: `[{"propertyId":"x","property":{"name":"a","value":"b"}}]`, Expected: existingProp, Error: true},
		{Name: "wrong prop json", Input: `[{"propertyx":"r","resx":{"name":"a","value":"b"}}]`, Expected: existingProp, Error: true},
		{Name: "wrong prop value json", Input: `[{"propertyId":"r","property":{"namex":"a","valuex":"b"}}]`, Expected: existingProp, Error: true},
		{Name: "invalid prop json", Input: `[{"propertyId":0,"property":{"name":"a","value":"b"}}]`, Expected: existingProp, Error: true},
		{Name: "invalid prop value json", Input: `[{"propertyId":"r","property":{"name":0,"value":0}}]`, Expected: existingProp, Error: true},
		{Name: "invalid store json", Input: `true`, Expected: existingProp, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[string, *testProperty]) {
		store := NewPropertyStore[testProperty]()
		store.Add("r", existingProp)

		err := json.Unmarshal([]byte(test.Input), store)
		hasErr := err != nil
		// fmt.Println(err)

		asserts.Equals(t, test.Error, hasErr)
		asserts.Equals(t, store.properties["r"], test.Expected)
		asserts.Equals(t, 1, len(store.properties))
	})
}
