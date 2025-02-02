package main

func main() {

	nums := []int{}
	m := make(map[int]bool)
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			m[nums[i]] = true
		} else {
			m[nums[i]] = false
		}
	}
	for j := 0; j < len(nums); j++ {
		if j%2 == 0 {
			for key, val := range m {
				if val == true {
					res = append(res, key)
					delete(m, key)
					break
				}

			}
		} else if j%2 != 0 {
			for key, val := range m {
				if val == false {
					res = append(res, key)
					delete(m, key)
					break
				}
			}

		}
	}
}
