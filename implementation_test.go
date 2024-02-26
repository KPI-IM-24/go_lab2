package lab2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluatePostfix(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   int
		wantErr    bool
	}{
		{"SimpleAddition", "0 4 +", 4, false},
		{"SimplePow", "3 4 ^", 81, false},
		{"SimpleExpression", "3 4 + 2 * 7 /", 2, false},
		{"EmptyString", "", 0, true},
		{"InvalidSymbol", "3 4 &", 0, true},
		{"OneOperand", "5", 5, false},
		{"ComplexExpression", "1 2 + 3 4 + *", 21, false},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := EvaluatePostfix(testCase.expression)

			if testCase.wantErr {
				assert.Error(t, err, "EvaluatePostfix() should return an error")
			} else {
				assert.NoError(t, err, "EvaluatePostfix() should not return an error")
				assert.Equal(t, testCase.expected, got, "evaluatePostfix() returned incorrect result")
			}
		})
	}
}
func ExampleEvaluatePostfix() {

	// Example 1: Simple addition
	result, err := EvaluatePostfix("3 4 +")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result)

	// Example 2: Complex expression
	result, err = EvaluatePostfix("5 1 2 + 4 * + 3 -")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result)

	// Example 3: Division by zero
	result, err = EvaluatePostfix("4 0 /")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Example 4: Unknown operator
	result, err = EvaluatePostfix("4 5 &")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Example 5: Power operation
	result, err = EvaluatePostfix("2 3 ^")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result)

	// Output:
	// 7
	// 14
	// Error: division by zero
	// Error: unknown operator: &
	// 8
}
