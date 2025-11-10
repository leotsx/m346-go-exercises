package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	messageEyes := fmt.Sprintln("the dice shows", eyes, "eyes")
	fmt.Fprint(os.Stdout, messageEyes)

	messageLog := fmt.Sprintln("the dice was rolled at", when)
	fmt.Fprint(os.Stdout, messageLog)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go TODO
}
