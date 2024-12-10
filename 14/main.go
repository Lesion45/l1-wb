package main

import (
	"fmt"
)

func customTypeOf(variable interface{}) string {
	switch variable.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case chan interface{}:
		return "chan interface{}"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func main() {
	var test1 string = "test"
	var test2 int = 99
	test3 := make(chan interface{})
	var test4 bool = true
	fmt.Println(customTypeOf(test1))
	fmt.Println(customTypeOf(test2))
	fmt.Println(customTypeOf(test3))
	fmt.Println(customTypeOf(test4))
}
