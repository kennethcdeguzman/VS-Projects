package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "this is a string!"
	string2 := "This is a string!"

	if strings.Compare(string1, string2) == 0 {
		fmt.Println("The strings are identical!")
	} else {
		fmt.Println("The strings do not match!")
	}
	if compareCaseIns(string1, string2) {
		fmt.Println("Strings are the same")
	}
}

func compareCaseIns(a, b string) bool {
	if len(a) == len(b) {
		if strings.Compare(strings.ToLower(a), strings.ToLower(b)) == 0 {
			return true
		}
	}
	return false
}
