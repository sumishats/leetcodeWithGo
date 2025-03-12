package main

import "fmt"

func main() {
	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	count := 0
	for _, arr := range grid {
		for _, val := range arr {
			if val < 0 {
				count++
			}
		}
	}
	fmt.Println(count)

}
