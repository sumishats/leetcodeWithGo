package main

import "fmt"

func main() {
	salary := []int{400, 3000, 1000, 2000}
	result := average(salary)
	fmt.Println(result)
}
func average(salary []int) float64 {
	sum := 0
	for i := 0; i < len(salary); i++ {
		for j := i; j < len(salary); j++ {
			if salary[i] > salary[j] {
				salary[i], salary[j] = salary[j], salary[i]
			}
		}
	}
	for k := 1; k < len(salary)-1; k++ {

		sum += salary[k]
	}
	return float64(sum) / float64(len(salary)-2)
}
