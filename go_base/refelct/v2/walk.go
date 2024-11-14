package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	for i := range val.NumField() {
		v := val.Field(i)
		if v.Kind() == reflect.String {
			fn(v.String())
		}
	}
}