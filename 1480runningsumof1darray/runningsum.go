package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := runningSum(nums)
	fmt.Println(sum)

}
func runningSum(nums []int) []int {

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums

	// sum := 0
	// arr := []int{}
	// for i := 0; i < len(nums); i++ {
	// 	sum += nums[i]
	// 	arr = append(arr, sum)
	// }
	// return arr
}
