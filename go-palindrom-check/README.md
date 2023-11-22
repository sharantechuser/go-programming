## Palindrom check

Function  `isPalindrome` accept `str` as string input parameter and return TRUE if str is a palindrome else return FALSE

````
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

````
