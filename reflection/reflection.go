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

	walkUtil := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkUtil(val.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkUtil(val.Field(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkUtil(val.MapIndex(key))
		}
	case reflect.String:
		fn(val.String())
	}

}
