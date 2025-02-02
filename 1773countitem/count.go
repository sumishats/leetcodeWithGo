package main

import "fmt"

func main() {

	var ruleKey string
	var ruleValue string
	items := [][]string{}
	count := 0
	store := make(map[string]int)
	store["type"] = 0
	store["color"] = 1
	store["name"] = 2

	for _, val := range items {
		if val[store[ruleKey]] == ruleValue {

			count++

		}

	}
	fmt.Println(count)
	// return count

}
