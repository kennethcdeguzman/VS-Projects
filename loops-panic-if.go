package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {

	//** for loop **//

	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			println("continuing...")
			continue
		}
	}

	//** for loop with a post clause **//

	for k := 0; k < 5; k++ {
		println(k)
		if k == 3 {
			println("to continue...")
			continue
		}
	}

	//** infinite loops **//
	var j int
	for {
		if j == 5 {
			break
		}
		println(j)
		j++
	}

	//** Loop over a collection **//
	slice := []int{1, 2, 3}
	for l, m := range slice {
		println(l, m)
	}
	names := map[int]string{1: "kenneth", 2: "mac", 3: "sprakatak"}
	for _, v := range names {
		println(v)
	}

	//** Panic **//

	println("Starting web server...")

	//panic("Something bad just happened")

	println("Web Server started")

	// ** If statements **//

	u1 := User{
		ID:        1,
		FirstName: "Kenneth",
		LastName:  "De Guzman",
	}

	u2 := User{
		ID:        2,
		FirstName: "Kenneth",
		LastName:  "Mac",
	}

	if u1 == u2 {
		println("The same user")
	} else if u1.FirstName == u2.FirstName {
		println("Same firstnames")
	} else {
		println("Different users")
	}
}

