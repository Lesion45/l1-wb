package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Start")
	Sleep(time.Second * 3)
	fmt.Println("End")
}
