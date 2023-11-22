package main

import "fmt"

func main() {

	input := "abcba"
	res := isPalindrome(input)
	if res {
		fmt.Println(input, "is a palindrome")
	} else {
		fmt.Println(input, "is not a palindrome")
	}

}

// return TRUE if a palindrome. else return FALSE
func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		if i > len(str)-1-i {
			break
		}
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
