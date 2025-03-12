package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 3
	store := fizzBuzz(n)
	fmt.Println(store)
}
func fizzBuzz(n int) []string {
	new := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			new = append(new, "FizzBuzz")
		} else if i%3 == 0 {
			new = append(new, "Fizz")

		} else if i%5 == 0 {
			new = append(new, "Buzz")
		} else {
			new = append(new, strconv.Itoa(i))
		}
	}
	return new
}
