package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	result := pivotIndex(nums)
	fmt.Println(result)

}
func pivotIndex(nums []int) int {
	leftsum := 0
	rightsum := 0
	for i := 0; i < len(nums); i++ {

		for j := 0; j < i; j++ {
			leftsum += nums[j]
		}
		for k := i + 1; k < len(nums); k++ {
			rightsum += nums[k]
		}

		if leftsum == rightsum {
			return i
		}
		leftsum = 0
		rightsum = 0
	}
	return -1
}

// 	leftsum := 0
// 	rightsum := 0

// 	for i := 0; i < len(nums); i++ {

// 		for j := 0; j < i; j++ {
// 			leftsum += nums[j]
// 		}

// 		for k := i + 1; k < len(nums); k++ {
// 			rightsum += nums[k]
// 		}

// 		if leftsum == rightsum {
// 			return i
// 		}

// 	}
// 	leftsum = 0
// 	rightsum = 0

// 	return -1
// }
