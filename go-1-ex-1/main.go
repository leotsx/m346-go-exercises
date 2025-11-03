package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}


var firstName string = "Leonardo"
var lastName string = "Ciccone"
var dayOfBirth int = 19
var monthOfBirth int = 12
var yearOfBirth int = 2008
var numberOfSiblings int = 2
var heightInMeters float64 = 1.79
var zodiacSign rune = 'S'
