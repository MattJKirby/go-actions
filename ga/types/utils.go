package types

import (
	"reflect"
	"strings"
)

func TypeName(i any) string {
	var t reflect.Type

	switch o := i.(type){
	case reflect.Type:
		t = o
	case reflect.Value:
		t = o.Type()
	default:
		t = reflect.TypeOf(o)
	}

	if t.Kind() == reflect.Pointer || t.Kind() == reflect.Slice {
		t = t.Elem()
	}

	if pkgPath := t.PkgPath(); pkgPath != "" {
		pkgPath = strings.TrimSuffix(pkgPath, "")
		return pkgPath + "/" + t.String()
	}

	return t.String()
}

func IsRefType(t reflect.Type) bool {
	switch t.Kind(){
	case reflect.Func, reflect.Interface:
		return true
	case reflect.Pointer:
		return t.Elem().Kind() == reflect.Struct
	default:
		return false
	}
}