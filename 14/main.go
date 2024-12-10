package main

import "fmt"

func customTypeOf(variable interface{}) string {
	switch variable.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case float32:
		return "float32"
	case float64:
		return "float64"
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
	var test3 float32 = 3.1
	var test4 float64 = 3.3
	var test5 int = 34
	var test6 bool = true
	fmt.Println(customTypeOf(test1))
	fmt.Println(customTypeOf(test2))
	fmt.Println(customTypeOf(test3))
	fmt.Println(customTypeOf(test4))
	fmt.Println(customTypeOf(test5))
	fmt.Println(customTypeOf(test6))
}
