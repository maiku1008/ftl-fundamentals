// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Add takes a list of numbers and returns the result of adding them together.
func Add(nums ...float64) float64 {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	default:
		var sum float64
		for _, n := range nums {
			sum += n
		}
		return sum
	}
}

// Subtract takes a list of numbers and returns the result of subtracting them from each other.
func Subtract(nums ...float64) float64 {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	default:
		var sum float64 = nums[0]
		for _, n := range nums[1:] {
			sum -= n
		}
		return sum
	}
}

// Multiply takes a list of numbers and multiplies them.
func Multiply(nums ...float64) float64 {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	default:
		var product float64 = nums[0]
		for _, n := range nums[1:] {
			product *= n
		}
		return product
	}
}

// Divide takes a list of numbers and divides them.
// Error is returned when dividing by 0.
func Divide(nums ...float64) (float64, error) {
	switch len(nums) {
	case 0:
		return 0, nil
	case 1:
		return nums[0], nil
	default:
		var product float64 = nums[0]
		for _, n := range nums[1:] {
			if n == 0 || product == 0 {
				return 0, fmt.Errorf("cannot divide by zero")
			}
			product /= n
		}
		return product, nil
	}
}

// Sqrt finds the square root of a number
// Running this on 0 or negative numbers returns an error
func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("not allowed to find the square root of a negative numbers")
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
	input = strings.ReplaceAll(input, " ", "")
	for _, op := range operations {
		if strings.Contains(input, op) {
			numbers := strings.Split(input, op)
			operation = op
			a, err = strconv.ParseFloat(numbers[0], 64)
			if err != nil {
				return 0, err
			}
			b, err = strconv.ParseFloat(numbers[1], 64)
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
		return 0, fmt.Errorf("no operator found in string")
	}
}
