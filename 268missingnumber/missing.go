package main

import "fmt"

func main() {

	nums := []int{3, 0, 1}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {

			if nums[j] == i {
				break
			}
			if nums[j] != i && j == len(nums)-1 {
				// return i
				fmt.Println(i)
			}
		}
	}
	// return -1
}
