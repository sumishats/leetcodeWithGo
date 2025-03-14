package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := intersection(nums1, nums2)
	fmt.Println(result)

}
func intersection(nums1 []int, nums2 []int) []int {
	new := []int{}
	store := make(map[int]int)

	for i := 0; i < len(nums1); i++ {

		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				store[nums1[i]]++
			}
		}
	}
	for k, _ := range store {
		new = append(new, k)
	}
	return new
}
