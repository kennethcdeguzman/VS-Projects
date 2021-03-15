package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	year := t.Year()
	month := t.Month()
	day := t.Day()

	fmt.Printf("today is %d/%d/%d\n", month, day, year)
	fmt.Println(t.Format(time.UnixDate))
	fmt.Println(t.Format(time.Kitchen))
	fmt.Println(t.Format("Mon Jan 2 15:04:05 MST 2006"))

	today := time.Now()
	future := today.AddDate(0, 6, 0)
	fmt.Printf("today is %v and futre date is %v", today, future)
}
