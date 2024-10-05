package utils

import (
	"fmt"
	"go-actions/ga/cr"

	"reflect"
	"testing"
)

func TestGetType(t *testing.T){
	tests := []cr.TestCase[any, reflect.Type]{
		{Name: "type string", Input: "string", Expected: reflect.TypeOf("string")},
		{Name: "type int", Input: 10, Expected: reflect.TypeOf(10)},
		{Name: "type value int", Input: reflect.ValueOf(10), Expected: reflect.TypeOf(10)},
		{Name: "type type int", Input: reflect.TypeOf(10), Expected: reflect.TypeOf(10)},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[any, reflect.Type]) {
		actual := GetType(test.Input)
		if actual != test.Expected {
			t.Errorf("test %s: got %q, expected %q", test.Name, actual, test.Expected)
		}
	})
}

func TestIsRefType(t *testing.T){
	tests := []cr.TestCase[reflect.Type, bool]{
		{Name: "func", Input: reflect.TypeOf(func() {}), Expected: true},
		{Name: "interface (stringer)", Input: reflect.TypeOf((*fmt.Stringer)(nil)).Elem(), Expected: true},
		{Name: "pointer to struct", Input: reflect.TypeOf(&struct{}{}), Expected: true},
		{Name: "pointer to non-struct", Input: reflect.TypeOf(new(int)), Expected: false},
		{Name: "struct", Input: reflect.TypeOf(struct{}{}), Expected: false},
		{Name: "non ref (int)", Input: reflect.TypeOf(10), Expected: false},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[reflect.Type, bool]) {
		actual := IsRefType(test.Input)
		if actual != test.Expected {
			t.Errorf("test %s: got %t, expected %t", test.Name, actual, test.Expected)
		}
	})
}

func TestGetValueType(t *testing.T){
	pointerToStuctType := reflect.TypeOf(&struct{}{})
	structType := reflect.TypeOf(struct{}{})
	nonStructType := reflect.TypeOf(10)
	intType := reflect.TypeOf(10)
	interfaceStringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	funcType := reflect.TypeOf(func() {})
	
	tests := []cr.TestCase[reflect.Type, reflect.Type]{
		{Name: "pointer to struct", Input: pointerToStuctType, Expected: structType},
		{Name: "struct", Input: structType, Expected: structType},
		{Name: "pointer to non-struct", Input: nonStructType, Expected: nonStructType},
		{Name: "int", Input: intType, Expected: intType},
		{Name: "interface (stringer)", Input: interfaceStringerType, Expected: interfaceStringerType},
		{Name: "func type", Input: funcType, Expected: funcType},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[reflect.Type, reflect.Type]) {
		actual := GetValueType(test.Input)
		if actual != test.Expected {
			t.Errorf("test %s: got %q, expected %q", test.Name, actual, test.Expected)
		}
	})
}

type someStruct struct {}

func TestGetTypeName(t *testing.T) {
	tests := []cr.TestCase[any, string]{
		{Name: "string", Input: "", Expected: "string"},
		{Name: "int", Input: 10, Expected: "int"},
		{Name: "struct", Input: someStruct{}, Expected: "someStruct"},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[any, string]) {
		actual := TypeName(test.Input)
		if actual != test.Expected {
			t.Errorf("test %s: got %q, expected %s", test.Name, actual, test.Expected)
		}
	})
}

func TestTypePath(t *testing.T){
	tests := []cr.TestCase[any, string]{
		{Name: "string path", Input: "someString", Expected: "/string"},
		{Name: "int path", Input: 42, Expected: "/int"},
		{Name: "interface path", Input: someStruct{}, Expected: "go-actions/ga/utils/utils.someStruct"},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[any, string]) {
		actual := TypePath(test.Input)
		if actual != test.Expected {
			t.Errorf("test %s: got %q, expected %s", test.Name, actual, test.Expected)
		}
	})
}