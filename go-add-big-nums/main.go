package main

import (
	"fmt"
	"strings"
)

func main() {

	result := add_large_numbers("301", "5")

	fmt.Println("Sum is ", result)
}

func add_large_numbers(num1, num2 string) string {

	var carry int
	var result string
	maxLen := max(len(num1), len(num2))

	num_1_padded := strings.Repeat("0", maxLen-len(num1)) + num1
	num_2_padded := strings.Repeat("0", maxLen-len(num2)) + num2

	for i := maxLen - 1; i >= 0; i-- {

		sum := int(num_1_padded[i]-'0') + int(num_2_padded[i]-'0') + carry
		carry = sum / 10
		result = fmt.Sprintf("%d", (sum%10)) + result

	}
	if carry > 0 {
		result = fmt.Sprintf("%d", carry) + result
	}

	return result
}
