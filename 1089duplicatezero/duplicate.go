package main

import "fmt"

func main() {

	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}
	fmt.Println(arr)
}
