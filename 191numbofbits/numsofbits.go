package main

import (
	"fmt"
)

func main() {
	n := 11
	v := hammingWeight(n)
	fmt.Println(v)

}

func hammingWeight(n int) int {

	count := 0
	binary := fmt.Sprintf("%b", n)

	for i := 0; i < len(binary); i++ {
		if string(binary[i]) == "1" {
			count++
		}
	}
	return count
}
