package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day byte
	Month byte
	Year int
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "Leonardo",
			LastName:  "Ciccone",
		},
		BirthDate: BirthDate{
			Day:   19,
			Month: 12,
			Year:  2008,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       '\u2650',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)

	fmt.Println("Siblings After:", me.NumberOfSiblings + 1)
}
