package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "I really ejoy the Go language. It's one of my favorites"
	searchTerm := "Go"
	result := strings.Contains(sampleString, searchTerm)

	fmt.Printf("the sample text includes %s: %t\n", searchTerm, result)

}
