package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name string
		a, b float64
		nums []float64
		want float64
	}{
		{name: "positive", a: 2, b: 2, want: 4},
		{name: "negative", a: -2, b: 5, want: 3},
		{name: "one zero", a: 0, b: 3, want: 3},
		{name: "two zeroes", a: 0, b: 0, want: 0},
		{name: "decimal", a: 2.5, b: 4, want: 6.5},
		{name: "many numbers", a: 1, b: 2, nums: []float64{3, 4, 5, 6, 7, 8, 9, 10}, want: 55},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b, tc.nums...)
		if tc.want != got {
			t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 1000; i++ {
		a, b := rand.Float64(), rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Errorf("adding %f, %f: want %f, got %f", a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name string
		a, b float64
		nums []float64
		want float64
	}{
		{name: "positive", a: 2, b: 2, want: 0},
		{name: "negative", a: -2, b: 5, want: -7},
		{name: "one zero", a: 0, b: 3, want: -3},
		{name: "two zeroes", a: 0, b: 0, want: 0},
		{name: "decimal", a: 2.5, b: 4, want: -1.5},
		{name: "many numbers", a: 1, b: 2, nums: []float64{3, 4, 5, 6, 7, 8, 9, 10}, want: -53},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b, tc.nums...)
		if tc.want != got {
			t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name string
		a, b float64
		nums []float64
		want float64
	}{
		{name: "positive", a: 2, b: 2, want: 4},
		{name: "negative", a: -2, b: 5, want: -10},
		{name: "one zero", a: 0, b: 3, want: 0},
		{name: "two zeroes", a: 0, b: 0, want: 0},
		{name: "decimal", a: 2.5, b: 4, want: 10},
		{name: "many numbers", a: 1, b: 2, nums: []float64{3, 4, 5, 6, 7, 8, 9, 10}, want: 3628800},
		{name: "many of the same numbers", a: 2, b: 2, nums: []float64{2, 2, 2, 2, 2, 2, 2, 2}, want: 1024},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b, tc.nums...)
		if tc.want != got {
			t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestAllOperationsWithNoError(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name string
		fn   func(a, b float64, nums ...float64) float64
		a, b float64
		want float64
	}{
		{name: "Addition", fn: calculator.Add, a: 2, b: 2, want: 4},
		{name: "Subtraction", fn: calculator.Subtract, a: 2, b: 2, want: 0},
		{name: "Multiplication", fn: calculator.Multiply, a: 2, b: 2, want: 4},
	}
	for _, tc := range testCases {
		got := tc.fn(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}

// closeEnough helps us confidently assert floating point values that have very small imprecision.
func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDivide(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		a, b        float64
		nums        []float64
		want        float64
		errExpected bool
	}{
		{name: "positive", a: 2, b: 2, want: 1, errExpected: false},
		{name: "negative", a: -4, b: 2, want: -2, errExpected: false},
		{name: "one zero", a: 0, b: 3, want: 0, errExpected: false},
		{name: "two zeroes", a: 0, b: 0, want: 0, errExpected: true},
		{name: "decimal", a: 2.5, b: 4, want: 0.625, errExpected: false},
		{name: "decimal with tiny floating point imprecision", a: 2, b: 3, want: 0.666667, errExpected: false},
		{name: "many numbers", a: 1024, b: 2, nums: []float64{2, 2, 2, 2, 2, 2, 2, 2}, want: 2, errExpected: false},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b, tc.nums...)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("test name: %s unexpected error status, got: %v", tc.name, err)
		}
		if !tc.errExpected && !closeEnough(tc.want, got, 0.001) {
			t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		num, want   float64
		errExpected bool
	}{
		{"positive", 4, 2, false},
		{"negative", -25, 5, true},
		{"zero", 0, 0, false},
		{"decimal", 30.25, 5.5, false},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.num)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("test name: %s unexpected error status, got: %v", tc.name, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestStringMath(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		input       string
		want        float64
		errExpected bool
	}{
		{"multiply", "2*2", 4, false},
		{"add", "1 + 1.5", 2.5, false},
		{"divide", "18   /   6", 3, false},
		{"subtract", "100-0.1", 99.9, false},
		{"bogus input", "foo_bar_baz", 0, true},
		{"return an error passed from other function", "1/0", 0, true},
	}
	for _, tc := range testCases {
		got, err := calculator.StringMath(tc.input)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("test name: %s unexpected error status, got: %v", tc.name, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}
