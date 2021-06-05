package main

import (
	"testing"
)

func TestQuadraticEquationIfBelowZero(t *testing.T) {
	QuadraticEquation(1, 2, 3)
	testCases := []struct {
		desc string
	}{
		{
			desc: "Testing if Quadric Equation is below zero",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			//assert := assert.New(t)

			//assert.Equal(t, result, "The two roots should be the same.")

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
