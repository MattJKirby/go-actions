// package main

// import (
// 	"go-actions/ga/action/instance/parameter"
// )

// func main() {

// 	// ex, _ := ga.GetActionDefinition(examples.ExampleAction{})
// 	// fmt.Println(ex)

// 	// action, _ := ga.GetAction(examples.ExampleAction{})

// 	// fmt.Println(action)

// 	// def := action.GetDef()

// 	// def.Execute()
// 	// def.IntegerParameter.SetValue(20)
// 	// def.Execute()

// 	s := parameter.NewStore()
// 	parameter.GetOrDefault("testInt", 10)(s)

// 	gotIntParam := parameter.GetOrDefault("testInt", 0)

// }

package main

import (
	"errors"
	"fmt"
	"reflect"
)

// TypedValue wraps a value with its type information
type TypedValue struct {
	Value any
	Type  reflect.Type
}

// TypedMap stores dynamically typed values with metadata
type TypedMap struct {
	data map[string]TypedValue
}

// NewTypedMap creates a new TypedMap
func NewTypedMap() *TypedMap {
	return &TypedMap{data: make(map[string]TypedValue)}
}

// Set stores a value along with its type information
func (tm *TypedMap) Set(key string, value any) {
	tm.data[key] = TypedValue{
		Value: value,
		Type:  reflect.TypeOf(value),
	}
}

// Get retrieves a value by key without requiring its type upfront
func (tm *TypedMap) Get(key string) (any, reflect.Type, error) {
	typedValue, exists := tm.data[key]
	if !exists {
		return nil, nil, errors.New("key does not exist")
	}
	return typedValue.Value, typedValue.Type, nil
}

func main() {
	tm := NewTypedMap()

	// Storing values of different types
	tm.Set("intValue", 42)
	tm.Set("stringValue", "hello")
	tm.Set("floatValue", 3.14)

	// Retrieving values dynamically
	value, valueType, err := tm.Get("intValue")
	if err != nil {
		fmt.Println("Error:", err)
	}
	val := reflect.New(valueType).Elem()

		fmt.Printf("Value: %v, Type: %v, instantiated: %v\n", value, valueType, val)
	
	
}
