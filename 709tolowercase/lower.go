package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "HELLO"
	result := toLowerCase(s)
	fmt.Println(result)

}
func toLowerCase(s string) string {
	store := ""
	store = strings.ToLower(s)

	return store
}
