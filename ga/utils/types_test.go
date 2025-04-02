package utils

import (
	"fmt"

	"reflect"
	"testing"
)

func TestGetType(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected reflect.Type
	}{
		{name: "type string", input: "string", expected: reflect.TypeOf("string")},
		{name: "type int", input: 10, expected: reflect.TypeOf(10)},
		{name: "type value int", input: reflect.ValueOf(10), expected: reflect.TypeOf(10)},
		{name: "type type int", input: reflect.TypeOf(10), expected: reflect.TypeOf(10)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := GetType(test.input)
			if actual != test.expected {
				t.Errorf("test %s: got %q, expected %q", test.name, actual, test.expected)
			}
		})
	}
}

func TestIsRefType(t *testing.T) {
	tests := []struct {
		name     string
		input    reflect.Type
		expected bool
	}{
		{name: "func", input: reflect.TypeOf(func() {}), expected: true},
		{name: "interface (stringer)", input: reflect.TypeOf((*fmt.Stringer)(nil)).Elem(), expected: true},
		{name: "pointer to struct", input: reflect.TypeOf(&struct{}{}), expected: true},
		{name: "pointer to non-struct", input: reflect.TypeOf(new(int)), expected: false},
		{name: "struct", input: reflect.TypeOf(struct{}{}), expected: false},
		{name: "non ref (int)", input: reflect.TypeOf(10), expected: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsRefType(test.input)
			if actual != test.expected {
				t.Errorf("test %s: got %t, expected %t", test.name, actual, test.expected)
			}
		})
	}
}

func TestGetValueType(t *testing.T) {
	pointerToStuctType := reflect.TypeOf(&struct{}{})
	structType := reflect.TypeOf(struct{}{})
	nonStructType := reflect.TypeOf(10)
	intType := reflect.TypeOf(10)
	interfaceStringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	funcType := reflect.TypeOf(func() {})

	tests := []struct {
		name     string
		input    reflect.Type
		expected reflect.Type
	}{
		{name: "pointer to struct", input: pointerToStuctType, expected: structType},
		{name: "struct", input: structType, expected: structType},
		{name: "pointer to non-struct", input: nonStructType, expected: nonStructType},
		{name: "int", input: intType, expected: intType},
		{name: "interface (stringer)", input: interfaceStringerType, expected: interfaceStringerType},
		{name: "func type", input: funcType, expected: funcType},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := GetValueType(test.input)
			if actual != test.expected {
				t.Errorf("test %s: got %q, expected %q", test.name, actual, test.expected)
			}
		})
	}
}

type someStruct struct{}

func TestGetTypeName(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{name: "string", input: "", expected: "string"},
		{name: "int", input: 10, expected: "int"},
		{name: "struct", input: someStruct{}, expected: "someStruct"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := TypeName(test.input)
			if actual != test.expected {
				t.Errorf("test %s: got %q, expected %s", test.name, actual, test.expected)
			}
		})
	}
}

func TestTypePath(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{name: "string path", input: "someString", expected: "/string"},
		{name: "int path", input: 42, expected: "/int"},
		{name: "interface path", input: someStruct{}, expected: "go-actions/ga/utils/utils.someStruct"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := TypePath(test.input)
			if actual != test.expected {
				t.Errorf("test %s: got %q, expected %s", test.name, actual, test.expected)
			}
		})
	}
}
