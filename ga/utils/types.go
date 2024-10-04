package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func GetType(i any) reflect.Type {
	switch t := i.(type) {
	case reflect.Type:
		return t
	case reflect.Value:
		return t.Type()
	default:
		return reflect.TypeOf(t)
	}
}

func GetValueType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Func || t.Kind() == reflect.Interface {
		return t 
	}

	if IsRefType(t) {
		return t.Elem()
	}
	return t
}

func TypePath(i any) string {
	t := GetType(i)
	valueType := GetValueType(t)

	return fmt.Sprintf("%s/%s", valueType.PkgPath(), valueType.String())
}

func TypeName(i any) string {
	s := strings.Split(GetType(i).String(), ".")
	return s[len(s)-1]
}

func IsRefType(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Func, reflect.Interface:
		return true
	case reflect.Pointer:
		return t.Elem().Kind() == reflect.Struct
	default:
		return false
	}
}
