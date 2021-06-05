package main

import (
	"fmt"
	"log"
	"math"
)

var a, b, c float64
var check string

func main() {
	fmt.Println("\nEnter a, b, c in sequence, separated by a space: ")
	fmt.Scanln(&a, &b, &c)
	Scan(a, b, c)
	fmt.Println("\nContinue?\nPress the 'Y' key if you want to continue, or any other key to exit.")
	fmt.Scanln(&check)
	Ask(check)
}

func Scan(a float64, b float64, c float64) (answer [3]float64) {

	answer = [3]float64{a, b, c}
	QuadraticEquation(a, b, c)
	return answer
}

func Ask(check string) (checked string) {

	a, b, c = 1, 2, 3
	checked = check
	if check == "Y" {
		main()
	} else {
		fmt.Println("Exiting...")
	}

	return checked
}

func QuadraticEquation(a float64, b float64, c float64) (result [2]float64) {

	var root1, root2, imaginary, discriminant float64

	if a == 0 {
		log.Fatalf("math: a = %v", a)
	}

	if b == 0 {
		log.Fatalf("math: b = %v", b)
	}

	if c == 0 {
		log.Fatalf("math: c = %v", c)
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

	}

	fmt.Printf("The result: %v\n", result)
	return result
}
