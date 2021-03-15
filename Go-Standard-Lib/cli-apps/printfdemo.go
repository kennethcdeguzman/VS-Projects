package main

import "fmt"

func main() {
	var age int = 10
	var name string = "Jeremy"

	/*
		%v - formats th value in a default format
		%s - fromats string values
		%d- formats decimal integers
		%g - formats the floating-point numbers
		%b - formats vbase 22 numbers
		%o - formats base 88 numbers
		%t - formats true or false values
	*/

	fmt.Printf("My name is %s and I am %d years old\n", name, age)

	var pie float32 = 3.141592

	fmt.Printf("pi is %.3f\n", pie)
	fmt.Printf("%d|%d|%d\n", 1, 2, 3)
}
