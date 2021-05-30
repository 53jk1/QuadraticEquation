package main

import (
	"testing"
)

func TestQuadraticEquationIfBelowZero(t *testing.T) {
	QuadraticEquation(1, 2, 3)
}

func TestQuadraticEquationIfGreaterThanZero(t *testing.T) {
	QuadraticEquation(1, 3, -4)
}

func TestQuadraticEquationIfZeroA(t *testing.T) {
	QuadraticEquation(0, 1, 4)
}

func TestQuadraticEquationIfZeroB(t *testing.T) {
	QuadraticEquation(1, 0, 4)
}
