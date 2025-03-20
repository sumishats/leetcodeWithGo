package main

func main() {

	nums := []int{1, 3, 9, 7, 6, 2, 4}
	store := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			if nums[i]%3 == 0 {
				store += nums[i]
			}
		}
	}
	// return store/2
}
