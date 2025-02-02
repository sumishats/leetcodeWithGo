package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 4}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]%2 == 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
}
