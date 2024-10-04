package utils

import (
	"reflect"
	"testing"
)

type GetTypeTest struct {
	name string
	input any
	expected reflect.Type
}

func TestGetType(t *testing.T){
	tests := []GetTypeTest{
		{"test string", "string", reflect.TypeOf("string")},
		{"test int", 10, reflect.TypeOf(10)},
		{"test value int", reflect.ValueOf(10), reflect.TypeOf(10)},
		{"test type int", reflect.TypeOf(10), reflect.TypeOf(10)},
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