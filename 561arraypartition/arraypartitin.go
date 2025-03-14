package main

import "fmt"

func main() {
	nums := []int{1, 4, 3, 2}
	result := arrayPairSum(nums)
	fmt.Println(result)
}
func arrayPairSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	for k := 0; k < len(nums); k++ {
		if k%2 != 0 {
			if nums[k-1] < nums[k] {
				sum += nums[k-1]
			} else {
				sum += nums[k]
			}
		}
	}
	return sum
}
