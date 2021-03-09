// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// Add takes a list of numbers and returns the result of adding them together.
func Add(a, b float64, nums ...float64) float64 {
	result := a + b
	for _, n := range nums {
		result += n
	}
	return result
}

// Subtract takes a list of numbers and returns the result of subtracting them from each other.
func Subtract(a, b float64, nums ...float64) float64 {
	result := a - b
	for _, n := range nums {
		result -= n
	}
	return result
}

// Multiply takes a list of numbers and multiplies them.
func Multiply(a, b float64, nums ...float64) float64 {
	result := a * b
	for _, n := range nums {
		result *= n
	}
	return result
}

// Divide takes a list of numbers and divides them.
// Error is returned when dividing by 0.
func Divide(a, b float64, nums ...float64) (float64, error) {
	err := errors.New("cannot divide by zero")
	if a == 0 || b == 0 {
		return 0, err
	}

	var result float64 = a / b
	for _, n := range nums {
		if n == 0 || result == 0 {
			return 0, err
		}
		result /= n
	}
	return result, nil
}

// Sqrt finds the square root of a number
// Running this on 0 or negative numbers returns an error
func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("not allowed to find the square root of a negative number: %f", num)
	}
	return math.Sqrt(num), nil
}

// StringMath takes a string, parses it and runs the math operation found therein
// Returns an error if it cannot do so
func StringMath(input string) (float64, error) {
	var a, b float64
	var operation string
	var err error

	operations := []string{"+", "-", "*", "/"}
	for _, op := range operations {
		if strings.Contains(input, op) {
			operation = op
			input = strings.Replace(input, op, " ", 1)
			_, err = fmt.Sscanf(input, "%f%f", &a, &b)
			if err != nil {
				return 0, err
			}
			break
		}
	}

	switch operation {
	case "+":
		return Add(a, b), nil
	case "-":
		return Subtract(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		return Divide(a, b)
	default:
		return 0, errors.New("no operator found in string")
	}
}
