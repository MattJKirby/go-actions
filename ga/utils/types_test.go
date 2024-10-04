package utils

import (
	"fmt"
	"reflect"
	"testing"
)

type getTypeTest struct {
	name string
	input any
	expected reflect.Type
}

func TestGetType(t *testing.T){
	tests := []getTypeTest{
		{"type string", "string", reflect.TypeOf("string")},
		{"type int", 10, reflect.TypeOf(10)},
		{"type value int", reflect.ValueOf(10), reflect.TypeOf(10)},
		{"type type int", reflect.TypeOf(10), reflect.TypeOf(10)},
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

type isRefTypeTest struct {
	name string
	input reflect.Type
	expected bool
}

func TestIsRefType(t *testing.T){
	tests := []isRefTypeTest{
		{"func", reflect.TypeOf(func() {}), true},
		{"interface (stringer)", reflect.TypeOf((*fmt.Stringer)(nil)).Elem(), true},
		{"pointer to struct", reflect.TypeOf(&struct{}{}), true},
		{"pointer to non-struct", reflect.TypeOf(new(int)), false},
		{"struct", reflect.TypeOf(struct{}{}), false},
		{"non ref (int)", reflect.TypeOf(10), false},
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

type getValueTypeTest struct {
	name string
	input reflect.Type
	expected reflect.Type
}

func TestGetValueType(t *testing.T){
	pointerToStuctType := reflect.TypeOf(&struct{}{})
	structType := reflect.TypeOf(struct{}{})
	nonStructType := reflect.TypeOf(10)
	intType := reflect.TypeOf(10)
	interfaceStringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	funcType := reflect.TypeOf(func() {})
	
	tests := []getValueTypeTest{
		{"pointer to struct", pointerToStuctType, structType},
		{"struct", structType, structType},
		{"pointer to non-struct", nonStructType, nonStructType},
		{"int", intType, intType},
		{"interface (stringer)", interfaceStringerType, interfaceStringerType},
		{"func type", funcType, funcType},
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

type typeNameTest struct {
	name string
	input any
	expected string
}

func TestGetTypeName(t *testing.T) {
	tests := []typeNameTest{
		{"string", "", "string"},
		{"int", 10, "int"},
		{"struct", typeNameTest{}, "typeNameTest"},
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

type typePathTest struct {
	name string
	input any
	expected string
}

func TestTypePath(t *testing.T){
	tests := []typePathTest{
		{"string path", "someString", "/string"},
		{"int path", 42, "/int"},
		{"interface path", typePathTest{}, "go-actions/ga/utils/utils.typePathTest"},
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