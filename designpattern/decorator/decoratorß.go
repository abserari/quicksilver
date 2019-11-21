package main

import (
	"fmt"
	"reflect"
)

func Decorate(impl interface{}) interface{} {
	fn := reflect.ValueOf(impl)

	inner := func(in []reflect.Value) []reflect.Value {
		f := reflect.ValueOf(impl)

		fmt.Println("Stuff before")
		// ...

		ret := f.Call(in)

		fmt.Println("Stuff after")
		// ...

		return ret
	}

	v := reflect.MakeFunc(fn.Type(), inner)

	return v.Interface()
}

var Add = Decorate(
	func(a, b int) int {
		return a + b
	},
).(func(a, b int) int)

func main() {
	fmt.Println(Add(1, 2))
}
