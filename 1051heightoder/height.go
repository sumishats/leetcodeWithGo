package main

import "fmt"

func main() {
	arr := []int{2, 4, 1, 6, 5}
	arr1 := make([]int, len(arr)) // Allocate space for arr1

	copy(arr1, arr)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
	fmt.Println(arr1)

}
