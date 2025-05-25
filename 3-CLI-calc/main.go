package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter expression (q to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "q" {
			break
		}

		tokens, err := tokenize(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		result, err := evaluate(tokens)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Result:", formatResult(result))
	}
}

func tokenize(input string) ([]string, error) {
	var tokens []string
	var buffer strings.Builder

	for _, r := range input {
		switch {
		case unicode.IsDigit(r) || r == '.':
			buffer.WriteRune(r)
		case r == '+' || r == '-' || r == '*' || r == '/' || r == '(' || r == ')':
			if buffer.Len() > 0 {
				tokens = append(tokens, buffer.String())
				buffer.Reset()
			}
			tokens = append(tokens, string(r))
		case unicode.IsSpace(r):
			if buffer.Len() > 0 {
				tokens = append(tokens, buffer.String())
				buffer.Reset()
			}
		default:
			return nil, fmt.Errorf("invalid character: %c", r)
		}
	}

	if buffer.Len() > 0 {
		tokens = append(tokens, buffer.String())
	}

	return tokens, nil
}

func evaluate(tokens []string) (float64, error) {
	postfix, err := infixToPostfix(tokens)
	if err != nil {
		return 0, err
	}

	stack := []float64{}

	for _, token := range postfix {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = a / b
			default:
				return 0, fmt.Errorf("unknown operator: %s", token)
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}

func infixToPostfix(tokens []string) ([]string, error) {
	var output []string
	var stack []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for _, token := range tokens {
		if _, err := strconv.ParseFloat(token, 64); err == nil {
			output = append(output, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1]
		} else {
			for len(stack) > 0 && stack[len(stack)-1] != "(" &&
				precedence[token] <= precedence[stack[len(stack)-1]] {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func formatResult(result float64) string {
	if result == math.Trunc(result) {
		return fmt.Sprintf("%.0f", result)
	}
	return fmt.Sprintf("%.2f", result)
}
