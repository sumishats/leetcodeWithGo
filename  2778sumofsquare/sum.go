package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	x := sumOfSquares(nums)
	fmt.Println(x)
}
func sumOfSquares(nums []int) int {
	v := 0
	for i := 0; i < len(nums); i++ {
		if len(nums)%(i+1) == 0 {
			v += nums[i] * nums[i]
		}
	}
	return v

}
