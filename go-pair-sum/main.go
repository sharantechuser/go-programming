package main

import "fmt"

func main() {

	input1 := []int{2, 5, -5, 10}
	input2 := 2

	// Print pair in optimal way O(nlon(n))
	mapInt := make(map[int]int)
	for index, v := range input1 {
		mapInt[v] = index
	}

	for index, out1 := range input1 {
		out2 := input2 - out1
		if idx, OK := mapInt[out2]; OK && idx != index {
			fmt.Println(out1, out2)
			break
		}
	}

}
