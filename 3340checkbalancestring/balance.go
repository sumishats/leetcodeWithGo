package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "1234"
	result := isBalanced(num)
	fmt.Println(result)
}
func isBalanced(num string) bool {
	oddsum := 0
	evensum := 0
	for i := 0; i < len(num); i++ {

		if i%2 == 0 {
			s, _ := strconv.Atoi(string(num[i]))
			evensum += s
		} else {
			s, _ := strconv.Atoi(string(num[i]))
			oddsum += s
		}
	}
	if oddsum == evensum {
		return true
	}
	return false
}
