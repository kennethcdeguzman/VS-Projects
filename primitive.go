package main

import (
	"fmt"
)

const (
	first  = iota
	second = 2 << iota
	fourth
)

const (
	fifth = iota
)

func main() {
	var firstname *string = new(string)
	*firstname = "Kenneth"
	fmt.Println(*firstname)

	surname := "De Guzman"
	ptr := &surname
	fmt.Println(ptr, *ptr)

	const i int = 3
	fmt.Println(float32(i) + 2.1)

	const j = 4
	fmt.Println(j + 3.3)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(fourth)
	fmt.Println(fifth)
}
