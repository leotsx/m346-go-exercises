package main
import "math"
import "fmt"
// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	hypotenuse := computeHypotenuse(3, 4)
	fmt.Printf("Hypotenuse: %.2f\n", hypotenuse)
}