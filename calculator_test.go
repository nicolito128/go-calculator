package calculator_test

import (
	"fmt"
	"math"
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

func TestResolveImplicitMult(t *testing.T) {
	{
		got, err := calculator.Resolve("(3)(4)")
		want := float64(12)

		if err != nil {
			t.Errorf("error occurred: %v", err)
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	{
		got, err := calculator.Resolve("(5)(5)2")
		want := float64(50)

		if err != nil {
			t.Errorf("error occurred: %v", err)
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestMathConstants(t *testing.T) {
	{
		got, err := calculator.Resolve("e")
		want := fmt.Sprintf("%f", math.E)

		if err != nil {
			t.Errorf("error occurred: %v", err)
		}

		s := fmt.Sprintf("%f", got)
		if s != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	{
		got, err := calculator.Resolve("pi")
		want := fmt.Sprintf("%f", math.Pi)

		if err != nil {
			t.Errorf("error occurred: %v", err)
		}

		s := fmt.Sprintf("%f", got)
		if s != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	{
		got, err := calculator.Resolve("phi")
		want := fmt.Sprintf("%f", math.Phi)

		if err != nil {
			t.Errorf("error occurred: %v", err)
		}

		s := fmt.Sprintf("%f", got)
		if s != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	{
		got, err := calculator.Resolve("ln10")
		want := fmt.Sprintf("%f", math.Ln10)

		if err != nil {
			t.Errorf("error occurred: %v", err)
		}

		s := fmt.Sprintf("%f", got)
		if s != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	{
		got, err := calculator.Resolve("ln2")
		want := fmt.Sprintf("%f", math.Ln2)

		if err != nil {
			t.Errorf("error occurred: %v", err)
		}

		s := fmt.Sprintf("%f", got)
		if s != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	{
		got, err := calculator.Resolve("pi/2 + (3)(4) - e")
		want := 10.852514498335852

		if err != nil {
			t.Errorf("error occurred: %v", err)
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestDivZero(t *testing.T) {
	_, err := calculator.Resolve("2 / 0")

	if err == nil {
		t.Errorf("Division by zero should return an error, but it didn't.")
	}
}
