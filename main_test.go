package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	a, b, c float64
	out     [2]float64
}{
	{1, 2, 3, [0.4 -2.4]},
	{2, 3, 4, },
}

func TestQuadraticEquationIfBelowZero(t *testing.T) {

	for _, tC := range testCases {
		testname := fmt.Sprintf("%d", QuadraticEquation(a, b, c))
		t.Run(testname, func(t *testing.T) {
			ans := QuadraticEquation(tC.a, tC.b, tC.c)
			if ans != tC.out {
				t.Errorf("got %q, want %q", ans, tC.out)
			}

		})
	}
}

func TestQuadraticEquationIfGreaterThanZero(t *testing.T) {
	QuadraticEquation(1, 3, -4)
}

// func TestQuadraticEquationIfZeroA(t *testing.T) {
// 	QuadraticEquation(0, 1, 4)
// }

// func TestQuadraticEquationIfZeroB(t *testing.T) {
// 	QuadraticEquation(1, 0, 4)
// }

// func TestQuadraticEquationIfZeroC(t *testing.T) {
// 	QuadraticEquation(1, 3, 0)
// }
