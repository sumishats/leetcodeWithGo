package main

func main() {

	accounts := [][]int{{1, 34, 5}, {6, 3, 6}, {78, 2}}

	v := 0
	for _, arr := range accounts {

		sum := 0

		for j := 0; j < len(arr); j++ {

			sum += arr[j]

		}

		if v < sum {

			v = sum
		}
	}

}
