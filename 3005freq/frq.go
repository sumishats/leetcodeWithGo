package main

func main() {
	nums := []int{10, 29, 37, 477, 57}
	store := make(map[int]int)
	maxFreq := 0
	count := 0

	for _, val := range nums {
		store[val]++
	}

	for _, val := range store {
		if val > maxFreq {
			maxFreq = val
			count = val
		} else if val == maxFreq {
			count += val
		}
	}

	// return count
}
