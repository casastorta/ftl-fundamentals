// Package calculator provides a library for simple calculations in Go.
package calculator

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multipilies two numbers and returns the result of multiplying the second
// against the first
func Multiply(a, b float64) float64 {
	return a * b
}
