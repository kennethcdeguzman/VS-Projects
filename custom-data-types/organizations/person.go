package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) socialSecurityNumber {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Name struct {
	first string
	last  string
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Identifiable
}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

func NewPerson(firstname, lastname string, identifiable Identifiable) Person {
	return Person{
		Name: Name{
			first: firstname,
			last:  lastname,
		},
		Identifiable: identifiable,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.first, p.last)
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("Twitter handler must start with @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}

func (p *Person) ID() string {
	return p.Identifiable.ID()
}

// func (p *Person) SetFirstName(newFirstName string) {
// 	p.FirstName = newFirstName
// }

// func (p *Person) GetFirstName() string {
// 	return p.FirstName
// }
