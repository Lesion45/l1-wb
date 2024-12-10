package main

import "fmt"

func main() {
	var a int = 4
	var b int = 5

	fmt.Println(a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a, b)
}
