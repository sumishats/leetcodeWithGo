package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x <= -1 {
		return false
	}

	stringval := strconv.Itoa(x)

	for i, j := 0, len(stringval)-1; i < len(stringval)/2; i, j = i+1, j-1 {
		if string(stringval[i]) != string(stringval[j]) {
			return false
		}
	}
	return true

}
func main() {

	res := isPalindrome(1001)
	fmt.Println(res)
}
