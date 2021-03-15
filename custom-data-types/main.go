package main

import (
	"fmt"

	organization "example.com/custom/datatypes/organizations"
)

func main() {
	p := organization.NewPerson("Kenneth", "De Guzman", organization.NewSocialSecurityNumber("1233-456-78"))
	fmt.Println(p)
	fmt.Println(p.FullName())
	err := p.SetTwitterHandler("@jam_wilis")
	if err != nil {
		fmt.Printf("An error occured setting twitter handler %s\n", err.Error())
	}
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectUrl())
}
