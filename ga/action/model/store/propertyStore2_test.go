package store

// import (
// 	"encoding/json"
// 	"fmt"
// 	"go-actions/ga/cr"
// 	"go-actions/ga/cr/asserts"
// 	"testing"
// )

// type testProperty struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// func (tp testProperty) GetPropertyId() string {
// 	return tp.Name
// }

// func TestAdd(t *testing.T) {
// 	store := NewPropertyStore[testProperty]()

// 	store.Add(&testProperty{})
// 	asserts.Equals(t, 1, len(store.properties))
// }

// func TestGet(t *testing.T) {
// 	store := NewPropertyStore[testProperty]()
// 	prop := &testProperty{Name: "r"}
// 	store.Add(prop)

// 	t.Run("test valid", func(t *testing.T) {
// 		val, err := store.Get("r")
// 		asserts.Equals(t, err, nil)
// 		asserts.Equals(t, prop, val)
// 	})

// 	t.Run("test invalid", func(t *testing.T) {
// 		val, err := store.Get("x")
// 		asserts.Equals(t, true, err != nil)
// 		asserts.Equals(t, nil, val)
// 	})
// }

// func TestGetOrDefault(t *testing.T) {
// 	prop := &testProperty{"r", "v"}

// 	t.Run("test get path", func(t *testing.T) {
// 		store := NewPropertyStore[testProperty]()
// 		store.Add(prop)
// 		retrieved := store.GetOrDefault(prop.Name, func() *testProperty {
// 			return &testProperty{"r", "v"}
// 		})
// 		asserts.Equals(t, prop, retrieved)
// 	})

// 	t.Run("test default path", func(t *testing.T) {
// 		store := NewPropertyStore[testProperty]()
// 		retrieved := store.GetOrDefault("someProp", func() *testProperty {
// 			return &testProperty{"r", "v"}
// 		})
// 		asserts.Equals(t, prop, retrieved)
// 	})
// }

// func TestCustomMarshalling(t *testing.T) {
// 	store := NewPropertyStore[testProperty]()
// 	prop := &testProperty{"r", "v"}
// 	marshalledProp, _ := json.Marshal(prop)
// 	store.Add(prop)

// 	expected := fmt.Sprintf(`[%s]`, string(marshalledProp))
// 	marshalledStore, err := json.Marshal(store)
// 	asserts.Equals(t, nil, err)
// 	asserts.Equals(t, []byte(expected), marshalledStore)
// }

// func TestCustomUnmarshalling(t *testing.T) {

// 	existingProp := &testProperty{"id", "b"}

// 	tests := []cr.TestCase[string, *testProperty]{
// 		{Name: "valid json, existing prop", Input: `[{"name":"id","value":"new"}]`, Expected: &testProperty{"id", "new"}, Error: false},
// 		{Name: "valid json, non-existing prop id", Input: `[{"name":"badid","value":"new"}]`, Expected: existingProp, Error: true},
// 		{Name: "wrong prop json", Input: `[{"namex":"id","value":"new"}]`, Expected: existingProp, Error: true},
// 		{Name: "invalid prop json", Input: `[{name:"id"},"value":"new"}]`, Expected: existingProp, Error: true},
// 		{Name: "invalid store json", Input: `true`, Expected: existingProp, Error: true},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			store := NewPropertyStore[testProperty]()
// 			store.Add(existingProp)

// 			err := json.Unmarshal([]byte(test.Input), store)
// 			hasErr := err != nil

// 			asserts.Equals(t, test.Error, hasErr)
// 			asserts.Equals(t, test.Expected, store.properties["id"])
// 			asserts.Equals(t, 1, len(store.properties))
// 		})
// 	}
// }
