package main

import (
	"fmt"
)

func setBit(num int64, i int, flag bool) int64 {
	if flag {
		return num | (1 << i)
	}

	return num &^ (1 << i)
}

func main() {
	var num int64 = 42
	var index int = 2
	var flag bool = true

	fmt.Printf("До: %d (bin: %064b)\n", num, num)
	num = setBit(num, index, flag)
	fmt.Printf("После: %d (bin: %064b)\n", num, num)
}
