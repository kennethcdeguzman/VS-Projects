package main

import (
	"fmt"

	"example.com/custom/types/media"
)

func main() {
	fmt.Println("My Favorite movie is")
	myFav := media.Movie{}
	myFav.Title = "The Mask"
	fmt.Printf("My Favorite movie is %s", myFav.Title)

}
