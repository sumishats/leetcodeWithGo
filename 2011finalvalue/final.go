package main

import "fmt"

func main() {
	slic := []string{"--x", "x++", "--x", "++x"}
	v := finalValueAfterOperations(slic)
	fmt.Println(v)
}
func finalValueAfterOperations(operations []string) int {
	count := 0
	for _, val := range operations {
		if string(val[1]) == "+" {
			count++
		} else {
			count--
		}
	}
	return count
}
