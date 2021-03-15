package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "      this is our text       "
	fmt.Printf("%q\n", sampleString)
	newString := strings.TrimSpace(sampleString)
	fmt.Printf("%q\n", newString)
	newStringRight := strings.TrimRight(sampleString, " ")
	fmt.Printf("%q\n", newStringRight)
	website := "https://wwww.pluralsight.com"
	domainName := strings.TrimPrefix(website, "https://")
	fmt.Println(domainName)
}
