package main

import "fmt"

func calculate(s string) int {
	isOperand := func(o rune) bool {
		switch o {
		case '(', ')', '+', '-':
			return false
		default:
			return true
		}
	}

	operationStack := []rune{}
	eq := ""

	for _, v := range s {
		if isOperand(v) {
			if v >= '0' && v <= '9' {
				eq += string(v)
			}
		} else {
			if len(operationStack) > 0 && operationStack[len(operationStack)-1] != '(' && v != '(' {
				eq += string(operationStack[len(operationStack)-1])
				operationStack = operationStack[:len(operationStack)-1]
			}

			if v == ')' {
				for len(operationStack) > 0 && operationStack[len(operationStack)-1] != '(' {
					eq += string(operationStack[len(operationStack)-1])
					operationStack = operationStack[:len(operationStack)-1]
				}
				if len(operationStack) > 0 && operationStack[len(operationStack)-1] == '(' {
					operationStack = operationStack[:len(operationStack)-1] // Remove '('
				}
			} else {
				operationStack = append(operationStack, v)
			}
		}
	}

	// Add remaining operators from stack
	for len(operationStack) > 0 {
		eq += string(operationStack[len(operationStack)-1])
		operationStack = operationStack[:len(operationStack)-1]
	}

	return evaluatePostfix(eq)
}

func evaluatePostfix(postfix string) int {
	stack := []int{}

	for _, c := range postfix {
		if c >= '0' && c <= '9' {
			stack = append(stack, int(c-'0'))
		} else if c == '+' || c == '-' {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if c == '+' {
				stack = append(stack, a+b)
			} else {
				stack = append(stack, a-b)
			}
		}
	}

	return stack[len(stack)-1]
}

func main() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
