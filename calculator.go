package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var errEmptyOperation = errors.New("operation not entered")
var errParseOperation = errors.New("operation cannot be parsed to numeric values")
var errDivisionByZero = errors.New("division by zero")

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
			return 0, errDivisionByZero
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

func IsOperation(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '^'
}

// Resolve resolves the expression passed as parameter.
func Resolve(expression string) (float64, error) {
	if expression == "" {
		return 0, errEmptyOperation
	}
	expression = strings.ToLower(expression)
	expression = strings.ReplaceAll(expression, " ", "")
	expression = strings.ReplaceAll(expression, ",", ".")
	expression = strings.ReplaceAll(expression, ")(", ")*(")

	// constants
	expression = parseConstants(expression)

	var values []float64
	var operations []rune
	var negative bool

	tokens := []rune(expression)
	i := 0
	for i < len(tokens) {
		if Precedence(string(tokens[i])) == 0 && !IsDigit(tokens[i]) && tokens[i] != '(' && tokens[i] != ')' {
			return 0, errParseOperation
		}

		// Current token is an opening
		// brace, so push to operations stack
		switch {
		case tokens[i] == '(':
			if len(values) > 0 && len(operations) == 0 {
				operations = append(operations, '*')
			}

			operations = append(operations, tokens[i])
		case tokens[i] == ')':
			j := i + 1
			if j < len(tokens) {
				if !IsOperation(tokens[j]) || tokens[j] == '(' || IsDigit(tokens[j]) {
					operations = append(operations, '*')
				}
			}

			for len(operations) > 0 && operations[len(operations)-1] != '(' {
				if len(operations) > 0 {
					if len(values) <= 1 {
						return float64(0), errors.New("invalid expression")
					}
				}

				err := resolveOperation(&values, &operations)
				if err != nil {
					return 0, err
				}
			}
			// remove the last opening brace
			if len(operations) > 0 {
				pop(&operations)
			}
		case IsDigit(tokens[i]):
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
		case tokens[i] == 'i':
			k := i - 1
			if i == 0 ||
				tokens[k] == '*' ||
				tokens[k] == '/' ||
				tokens[k] == '+' ||
				tokens[k] == '-' ||
				tokens[k] == '^' ||
				tokens[k] == '(' ||
				tokens[k] == ')' {
				negative = true
				continue
			}
		default:
			for len(operations) > 0 && Precedence(string(operations[len(operations)-1])) >= Precedence(string(tokens[i])) {
				if len(operations) > 0 {
					if len(values) <= 1 {
						return float64(0), errors.New("invalid expression")
					}
				}

				err := resolveOperation(&values, &operations)
				if err != nil {
					return 0, err
				}
			}

			operations = append(operations, tokens[i])
		}

		if negative {
			topVal := values[len(values)-1]
			topVal *= (-1)
			values[len(values)-1] = topVal
			negative = false
		}

		i++
	}

	if len(operations) > 0 {
		if len(values) <= 1 {
			return float64(0), errors.New("invalid expression")
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

func parseConstants(exp string) string {
	// format
	e := strconv.FormatFloat(math.E, 'f', -1, 64)
	pi := strconv.FormatFloat(math.Pi, 'f', -1, 64)
	phi := strconv.FormatFloat(math.Phi, 'f', -1, 64)
	ln10 := strconv.FormatFloat(math.Ln10, 'f', -1, 64)
	ln2 := strconv.FormatFloat(math.Ln2, 'f', -1, 64)
	// replace
	exp = strings.ReplaceAll(exp, "e", e)
	exp = strings.ReplaceAll(exp, "pi", pi)
	exp = strings.ReplaceAll(exp, "phi", phi)
	exp = strings.ReplaceAll(exp, "ln10", ln10)
	exp = strings.ReplaceAll(exp, "ln2", ln2)
	return exp
}
