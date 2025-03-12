package main

import "fmt"

func main() {

	nums := []int{1, 2, 5, 2, 3}
	target := 2
	store := targetIndices(nums, target)
	fmt.Println(store)
}

func targetIndices(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		if nums[i] == target {
			result = append(result, i)
		}
	}
	return result

}
