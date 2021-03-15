package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleString := "This isa stirng. There are many strings like it, but this one is mine"
	r, _ := regexp.Compile(`s([a-z]+)g`)
	fmt.Println(r.MatchString(sampleString))
}
