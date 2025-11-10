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

	if err := os.WriteFile("eyes.txt", []byte(messageEyes), 0o644); err != nil {
		panic(err)
	}
	if err := os.WriteFile("dice.log", []byte(messageLog), 0o644); err != nil {
		panic(err)
	}
}
