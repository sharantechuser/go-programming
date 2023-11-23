
## Find pair of elements there sum should match the target value

The below  code accept two parameter `input1` as []int and `input2` as int target value. 

Below code will find the pair of elements from a slice (input1) to match there sum with target value(input2)

````

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

````