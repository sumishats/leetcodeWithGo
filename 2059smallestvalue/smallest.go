package main

import "fmt"

func main() {

	nums := []int{4, 3, 2, 1}
	store := smallestEqual(nums)
	fmt.Println(store)
}
func smallestEqual(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if i%10 == nums[i] {
			return i
		}
	}
	return -1
}
