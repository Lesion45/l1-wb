package main

import "fmt"

func main() {
	var temperature = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	var temperatureSet = make(map[int][]float32)

	for _, v := range temperature {
		temperatureSet[int(v/10)*10] = append(temperatureSet[int(v/10)*10], v)
	}

	fmt.Println(temperatureSet)
}
