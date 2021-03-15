package models

type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	employees []*Employee
	nextID2   = 1
)
