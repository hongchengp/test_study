package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := getValue(x)
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := range val.NumField() {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := range val.Len() {
			walk(val.Index(i).Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value{
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}