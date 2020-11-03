// Package reflection implements a function walk(x interface{}, fn func(string)) which takes a
// struct x and calls fn for all strings fields found inside. difficulty level: recursively.
package reflection

import (
	"reflect"
)

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	// slices
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	// struct fields
	var fieldCount int
	if val.Kind().String() == "struct" {
		fieldCount = val.NumField()
	}
	for i := 0; i < fieldCount; i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
