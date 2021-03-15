package main

import (
	"fmt"

	"example.com/custom/datatypes/media"
)

func main() {
	fmt.Println("My Favorite movie is ")

	favMovie := media.Movie{}
	favMovie.NewMovie("Matrix", media.PG, 50.13)

	fmt.Printf("My favorite movie is %s\n", favMovie.GetTitle())
	fmt.Printf("It was rated %s\n", favMovie.GetRating())
	fmt.Printf("It made %f Million in the box office\n", favMovie.GetBoxOffice())

	favMovie.SetTitle("Dumb and Dumber")
	fmt.Println(favMovie.GetTitle())
}
