package main

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	input := val.Field(0)
	fmt.Println(val)
	fmt.Println(input)
	fn(input.String())
}
