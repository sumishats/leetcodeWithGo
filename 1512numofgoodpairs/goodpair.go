package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	result := numIdenticalPairs(nums)
	fmt.Println(result)
}
func numIdenticalPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				count++
			}
		}
	}
	return count

}
