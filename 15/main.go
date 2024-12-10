package main

import "fmt"

func createHugeString(size int) string {
	return string(make([]byte, size))
}

var justString string

// При выполнении justString = v[:100] мы сохраним ссылку на срез строки.
// Это приведет к тому, что значение переменной v не будет очищено, так как на v будет ссылаться justString.
// Чтобы решить эту проблему, нужно создать копию подстроки, чтобы она больше не ссылалась на оригинальную строку.

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
