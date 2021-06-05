package main

import (
	"fmt"
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
