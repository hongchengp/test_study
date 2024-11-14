package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	in := val.Field(0)
	fn(in.String())
}