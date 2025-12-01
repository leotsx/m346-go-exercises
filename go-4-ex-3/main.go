package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a, b, c float64) []float64 {
	discriminant := math.Pow(b, 2) - 4*a*c
	if discriminant > 0 {
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return []float64{root1, root2}
	} else if discriminant == 0 {
		root := -b / (2 * a)
		return []float64{root}
	} else {
		return []float64{}
	}
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println(computeQuadraticFormula(3, 4, 1))
	fmt.Println(computeQuadraticFormula(2, 4, 2))
	fmt.Println(computeQuadraticFormula(3, 4, 2))

}
