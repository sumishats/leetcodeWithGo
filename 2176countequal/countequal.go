package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 2, 2, 1, 3}
	k := 2
	store := countPairs(nums, k)
	fmt.Println(store)
}
func countPairs(nums []int, k int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				count++
			}
		}
	}
	return count

}
