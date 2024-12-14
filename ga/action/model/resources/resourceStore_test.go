package resources

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/cr"
	"go-actions/ga/cr/asserts"
	"testing"
)

type testResource struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func TestAdd(t *testing.T) {
	store := NewResourceStore[testResource]()

	store.Add("r", &testResource{})
	asserts.Equals(t, 1, len(store.resources))
}

func TestGet(t *testing.T) {
	store := NewResourceStore[testResource]()
	resource := &testResource{}
	store.Add("r", resource)

	t.Run("test valid", func(t *testing.T) {
		val, err := store.Get("r")
		asserts.Equals(t, err, nil)
		asserts.Equals(t, resource, val)
	})

	t.Run("test invalid", func(t *testing.T) {
		val, err := store.Get("x")
		asserts.Equals(t, true, err != nil)
		asserts.Equals(t, nil, val)
	})
}

func TestGetOrDefault(t *testing.T) {
	resource := &testResource{"r", "v"}

	t.Run("test get path", func(t *testing.T) {
		store := NewResourceStore[testResource]()
		store.Add(resource.Name, resource)
		retrieved := store.GetOrDefault(resource.Name, func() *testResource {
			return &testResource{"r", "v"}
		})
		asserts.Equals(t, resource, retrieved)
	})

	t.Run("test default path", func(t *testing.T) {
		store := NewResourceStore[testResource]()
		retrieved := store.GetOrDefault("someResource", func() *testResource {
			return &testResource{"r", "v"}
		})
		asserts.Equals(t, resource, retrieved)
	})
}

func TestCustomMarshalling(t *testing.T) {
	store := NewResourceStore[testResource]()
	resource := &testResource{"r", "v"}
	marshalledResource, _ := json.Marshal(resource)
	store.Add("r", resource)

	expected := fmt.Sprintf(`[{"id":"r","attributes":%s}]`, string(marshalledResource))
	marshalledStore, err := json.Marshal(store)
	asserts.Equals(t, nil, err)
	asserts.Equals(t, []byte(expected), marshalledStore)
}

func TestCustomUnmarshalling(t *testing.T) {

	existingResource := &testResource{"r", "v"}

	tests := []cr.TestCase[string, *testResource]{
		{Name: "valid json, existing resource", Input: `[{"id":"r","attributes":{"name":"a","value":"b"}}]`, Expected: &testResource{"a", "b"}, Error: false},
		{Name: "valid json, non-existing resource id", Input: `[{"id":"x","attributes":{"name":"a","value":"b"}}]`, Expected: existingResource, Error: true},
		{Name: "wrong resource json", Input: `[{"idx":"r","resx":{"name":"a","value":"b"}}]`, Expected: existingResource, Error: true},
		{Name: "wrong resource value json", Input: `[{"id":"r","attributes":{"namex":"a","valuex":"b"}}]`, Expected: existingResource, Error: true},
		{Name: "invalid resource json", Input: `[{"id":0,"attributes":{"name":"a","value":"b"}}]`, Expected: existingResource, Error: true},
		{Name: "invalid resource value json", Input: `[{"id":"r","attributes":{"name":0,"value":0}}]`, Expected: existingResource, Error: true},
		{Name: "invalid store json", Input: `true`, Expected: existingResource, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[string, *testResource]) {
		store := NewResourceStore[testResource]()
		store.Add("r", existingResource)

		err := json.Unmarshal([]byte(test.Input), store)
		hasErr := err != nil
		// fmt.Println(err)

		asserts.Equals(t, test.Error, hasErr)
		asserts.Equals(t, store.resources["r"], test.Expected)
		asserts.Equals(t, 1, len(store.resources))
	})
}
