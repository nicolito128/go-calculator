package calculator_test

import (
	"testing"

	"github.com/nicolito128/go-calculator"
)

func TestResolve(t *testing.T) {
	got, err := calculator.Resolve("10 / 2 + 1")
	want := float64(6)

	if err != nil {
		t.Errorf("error occurred: %v", err)
	}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestResolveParenthesis(t *testing.T) {
	got, err := calculator.Resolve("(10 / 2) + (3 / 3)")
	want := float64(6)

	if err != nil {
		t.Errorf("error occurred: %v", err)
	}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDivZero(t *testing.T) {
	_, err := calculator.Resolve("2 / 0")

	if err == nil {
		t.Errorf("Division by zero should return an error, but it didn't.")
	}
}
