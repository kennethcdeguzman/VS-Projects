package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v", text)
	fmt.Printf("We are running at version %v in %v os", runtime.Version(), runtime.GOOS)
} 
