package calculator_test

import (
	"calculator"
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
		nums []float64
		want float64
	}{
		{"positive", []float64{2, 2}, 4},
		{"negative", []float64{-2, 5}, -10},
		{"one zero", []float64{0, 3}, 0},
		{"two zeroes", []float64{0, 0}, 0},
		{"decimal", []float64{2.5, 4}, 10},
		{"empty slice", []float64{}, 0},
		{"one item in slice", []float64{3}, 3},
		{"many numbers", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3628800},
		{"many of the same numbers", []float64{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 1024},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.nums...)
		if tc.want != got {
			t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		nums        []float64
		want        float64
		errExpected bool
	}{
		{"positive", []float64{2, 2}, 1, false},
		{"negative", []float64{-2, 5}, -0.4, false},
		{"one zero", []float64{0, 3}, 0, true},
		{"two zeroes", []float64{0, 0}, 0, true},
		{"decimal", []float64{2.5, 4}, 0.625, false},
		{"empty slice", []float64{}, 0, false},
		{"one item in slice", []float64{3}, 3, false},
		{"many numbers", []float64{1024, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 2, false},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.nums...)
		if tc.errExpected {
			if err == nil {
				t.Errorf("test name: %s expected error, got nil", tc.name)
			}
		} else {
			if err != nil {
				t.Errorf("test name: %s returned unexpected error: %s", tc.name, err)
			}
			if tc.want != got {
				t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
			}
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
		if tc.errExpected {
			if err == nil {
				t.Errorf("test name: %s expected error, got nil", tc.name)
			}
		} else {
			if err != nil {
				t.Errorf("test name: %s returned unexpected error: %s", tc.name, err)
			}
			if tc.want != got {
				t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
			}
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
		if tc.errExpected {
			if err == nil {
				t.Errorf("test name: %s expected error, got nil", tc.name)
			}
		} else {
			if err != nil {
				t.Errorf("test name: %s returned unexpected error: %s", tc.name, err)
			}
			if tc.want != got {
				t.Errorf("test name: %s, want %f, got %f", tc.name, tc.want, got)
			}
		}
	}
}
