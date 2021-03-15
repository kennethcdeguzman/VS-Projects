package main

import (
	"fmt"
)

func main() {

	/** ARRAYS **/

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[0])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println(arr2[0])

	/** SLICE **/

	slice := arr[:]
	fmt.Println(slice)

	slice2 := []int{4, 5, 6}
	fmt.Println(slice2)

	slice2 = append(slice2, 7)
	slice2 = append(slice2, 8)
	fmt.Println(slice2)

	s2 := slice2[1:]
	s3 := slice2[:2]
	fmt.Println(s2, s3)

	/** MAPS **/

	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 32
	fmt.Println(m)
	fmt.Println(m["foo"])
	delete(m, "foo")
	fmt.Println(m)

	/** STRUCTS **/

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Kenneth"
	u.LastName = "De Guzman"
	fmt.Println(u)

	u2 := user{ID: 1, FirstName: "Kenneth", LastName: "De Guzman"}

	fmt.Println(u2)
}
