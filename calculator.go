package calculator

import (
	"errors"
	"math"
	"strings"
)

var emptyOperationError = errors.New("Operation not entered!")
var convertError = errors.New("The operation cannot be converted to numeric values.")
var divisionByZero = errors.New("Division by zero!")

// Precedence returns the precedence of the operation.
// "+, -": 1; "*, /": 2; "^": 3
func Precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	}
	return 0
}

// ApplyOperation applies the operation to the values
func ApplyOperation(a, b float64, symbol rune) (float64, error) {
	switch symbol {
	case '+':
		return (a + b), nil
	case '-':
		return (a - b), nil
	case '*':
		return (a * b), nil
	case '/':
		if b == 0 {
			return 0, divisionByZero
		}

		return (a / b), nil
	case '^':
		return math.Pow(a, b), nil
	}

	return 0, nil
}

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// Resolve resolves the expression passed as parameter.
func Resolve(expression string) (float64, error) {
	if expression == "" {
		return 0, emptyOperationError
	}
	expression = strings.ReplaceAll(expression, " ", "")
	expression = strings.ReplaceAll(expression, ",", ".")

	var values []float64
	var operations []rune
	var negative bool

	tokens := []rune(expression)
	for i := 0; i < len(tokens); i++ {
		if Precedence(string(tokens[i])) == 0 && !IsDigit(tokens[i]) && tokens[i] != '(' && tokens[i] != ')' {
			return 0, convertError
		}

		// Current token is an opening
		// brace, so push to operations stack
		if tokens[i] == '(' {
			operations = append(operations, tokens[i])
		} else if IsDigit(tokens[i]) {
			// Contain the final number
			var val float64
			var existsDecimal bool
			var decimalPlaces int

			for i < len(tokens) && (IsDigit(tokens[i]) || tokens[i] == '.') {
				// If the current token is a number with decimal digits
				if tokens[i] == '.' {
					existsDecimal = true
					i++
					decimalPlaces++
					continue
				}

				if existsDecimal {
					digit := float64(tokens[i]-'0') / math.Pow(10, float64(decimalPlaces))

					// Last decimal digit check.
					if digit >= 0 {
						// Case: 11 + 0.1 = 11.1
						val += digit
						decimalPlaces++
					} else {
						val = val*10 + digit
						decimalPlaces++
					}
				} else {
					// Normal addition by digit.
					digit := float64(tokens[i] - '0')
					val = val*10 + digit
				}

				i++
			}

			values = append(values, val)
			i--
		} else if tokens[i] == ')' {
			for len(operations) > 0 && operations[len(operations)-1] != '(' {
				err := resolveOperation(&values, &operations)
				if err != nil {
					return 0, err
				}
			}

			// remove the last opening brace
			if len(operations) > 0 {
				pop(&operations)
			}
		} else {
			if tokens[i] == '-' && (i == 0 || tokens[i-1] == '*' || tokens[i-1] == '/' || tokens[i-1] == '+' || tokens[i-1] == '-' || tokens[i-1] == '^' || tokens[i-1] == '(' || tokens[i-1] == ')') {
				negative = true
				continue
			} else {
				for len(operations) > 0 && Precedence(string(operations[len(operations)-1])) >= Precedence(string(tokens[i])) {
					err := resolveOperation(&values, &operations)
					if err != nil {
						return 0, err
					}
				}

				operations = append(operations, tokens[i])
			}
		}

		if negative {
			topVal := values[len(values)-1]
			topVal *= (-1)
			values[len(values)-1] = topVal
			negative = false
		}
	}

	for len(operations) > 0 {
		err := resolveOperation(&values, &operations)
		if err != nil {
			return 0, err
		}
	}

	return pop(&values), nil
}

// resolveOperation take the last two values and the last operation
// then apply the operation to the values and return an error
func resolveOperation(values *[]float64, operations *[]rune) error {
	valB := pop(values)
	valA := pop(values)
	op := pop(operations)

	result, err := ApplyOperation(valA, valB, op)
	if err != nil {
		return err
	}

	*values = append(*values, result)
	return nil
}

// pop removes and returns the last element of the values slice.
func pop[T any](arr *[]T) T {
	// Get the last element
	top := (*arr)[len(*arr)-1]
	// Remove last element
	*arr = (*arr)[:len(*arr)-1]
	return top
}
