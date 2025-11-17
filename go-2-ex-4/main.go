package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Name     string
		Students []Student
	}
	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int][]Class{
		346: {
			{
				Name: "Ina24a",
				Students: []Student{
					{"Alice", "Smith"},
					{"Bob", "Johnson"},
					{"Charlie", "Brown"},
				},
			},
			{
				Name: "Ina24b",
				Students: []Student{
					{"David", "Lee"},
					{"Leonardo", "Ciccone"},
					{"Frank", "Miller"},
				},
			},
		},
		104: {
			{
				Name: "Ina24c",
				Students: []Student{
					{"Grace", "Wilson"},
					{"Henry", "Moore"},
					{"Ivy", "Taylor"},
				},
			},
		},
		117: {
			{
				Name: "Ina24d",
				Students: []Student{
					{"Jack", "Anderson"},
					{"Kate", "Thomas"},
					{"Ramon", "Zielke"},
				},
			},
			{
				Name: "Ina24e",
				Students: []Student{
					{"Mia", "Garcia"},
					{"Noah", "Martinez"},
					{"Olivia", "Rodriguez"},
				},
			},
		},
	}
	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
