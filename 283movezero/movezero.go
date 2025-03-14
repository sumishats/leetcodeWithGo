package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	for j := count; j < len(nums); j++ {
		nums[j] = 0
	}
	fmt.Println(nums)
}
