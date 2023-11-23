package main

import "fmt"

func main() {

	var n = 2
	res := printBalancedParenthesis(n)

	fmt.Println(res)

}

func printBalancedParenthesis(n int) []string {
	var res []string
	parenthesis(&res, n, 0, 0, "")

	return res
}

// Print balanced parenthesis
func parenthesis(res *[]string, n, open, close int, curr_str string) {

	if open == n && close == n {
		*res = append(*res, curr_str)
	}
	if open < n {
		parenthesis(res, n, open+1, close, curr_str+"{")
	}
	if open > close {
		parenthesis(res, n, open, close+1, curr_str+"}")

	}
}
