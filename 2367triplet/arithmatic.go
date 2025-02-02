package main

func main() {
	nums := []int{1, 3, 28, 937, 66, 2}
	diff := 3
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					count++
				}
			}
		}
	}
	// return count
}
