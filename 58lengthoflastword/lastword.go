package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "hello world "
	v := strings.TrimSpace(s)

	count := 0
	for i := len(v) - 1; i > 0; i-- {
		if string(v[i]) != " " {
			count++
		} else {
			break
		}

	}
	fmt.Println(count)

}
