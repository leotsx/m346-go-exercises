package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1] // 1,1,2,0,0
	fibs[3] = fibs[1] + fibs[2] // 1,1,2,3,0
	fibs[4] = fibs[2] + fibs[3] // 1,1,2,3,5

	fibs = append(fibs, 8)
	fibs = append(fibs, 13, 21, 34)
	fmt.Println(fibs) 
}
