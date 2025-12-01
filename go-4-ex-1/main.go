package main
import "fmt"

// TODO: implement the function computeGrade
func computeGrade(score float32, maxpoints float32) float32 {
	grade := (score / maxpoints) * 5 + 1
	return grade
}



func main() {
	// TODO: call the function computeGrade
	grade := computeGrade(85, 100)
	fmt.Printf("Grade: %.2f\n", grade)
}
