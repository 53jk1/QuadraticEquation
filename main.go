package main

import (
	"fmt"
	"log"
	"math"
)

var a, b, c float64
var check string

func main() {
	Scan()
}

func Scan() {
	fmt.Println("\nEnter a, b, c in sequence, separated by a space: ")
	fmt.Scanln(&a, &b, &c)
	QuadraticEquation(a, b, c)
	Ask()
}

func Ask() {
	fmt.Println("\nContinue?\nPress the 'Y' key if you want to continue, or any other key to exit.")
	fmt.Scanln(&check)
	if check == "Y" {
		Scan()
		fmt.Println()
	}

}

func QuadraticEquation(a float64, b float64, c float64) (result [2]float64) {

	var root1, root2, imaginary, discriminant float64

	if a == 0 {
		log.Fatalf("math: a = 0")
	}

	if b == 0 {
		log.Fatalf("math: b = 0")
	}

	if c == 0 {
		log.Fatalf("math: c = 0")
	}

	discriminant = (b * b) - (4 * a * c)

	switch {

	case discriminant > 0:
		result1 := (-b + math.Sqrt(discriminant)/(2*a))
		result2 := (-b - math.Sqrt(discriminant)/(2*a))

		result = [2]float64{result1, result2}

	case discriminant == 0:
		result1 := -b / (2 * a)
		result2 := -b / (2 * a)
		result = [2]float64{result1, result2}

	case discriminant < 0:
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		imaginary = math.Sqrt(-discriminant) / (2 * a)
		result1 := root1 + imaginary
		result2 := root2 - imaginary

		result = [2]float64{result1, result2}

	default:
		log.Fatalf("not valid number %v", result)

	}

	fmt.Printf("The result: %v\n", result)
	return result
}
