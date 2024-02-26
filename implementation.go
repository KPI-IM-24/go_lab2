package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type DefaultPostfixCalculator struct {
}

func (dpc *DefaultPostfixCalculator) EvaluatePostfix(expression string) (int, error) {

	var stack []int

	tokens := strings.Fields(expression)

	for _, token := range tokens {
		if value, err := strconv.Atoi(token); err == nil {
			stack = append(stack, value)
		} else {
			stackLen := len(stack)

			if stackLen < 2 {
				return 0, fmt.Errorf("invalid expression was used")
			}

			a, b := stack[stackLen-2], stack[stackLen-1]
			stack = stack[:stackLen-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				stack = append(stack, a/b)
			case "^":
				stack = append(stack, int(math.Pow(float64(a), float64(b))))
			default:
				return 0, fmt.Errorf("unknown operator: %s", token)
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}
