package main

import "fmt"

func main() {
	s := "hlomynameammu"
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	fmt.Println(result)
}
