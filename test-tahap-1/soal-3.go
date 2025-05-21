package main

import (
	"fmt"
)

func isMatchingPair(open, close rune) bool {
	switch open {
	case '<':
		return close == '>'
	case '(':
		return close == ')'
	case '[':
		return close == ']'
	case '{':
		return close == '}'
	}
	return false
}

func isValidBracketString(input string) bool {
	stack := []rune{}

	for _, char := range input {
		switch char {
		case '<', '(', '[', '{':
			stack = append(stack, char)
		case '>', ')', ']', '}':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if !isMatchingPair(top, char) {
				return false
			}

			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	// Example usage
	tests := []string{
		"{{[<>[{{}}]]}}",                     // true
		"{<{[[{{[]<{{[{[]<>}]}}<>>}}]]}>}",   // true
		"<<<<<<<[{[{[{[<<{}>>]}]}]}]>>>>>>>", // true
		"]",                                  // false
		"][",                                 // false
		"{[<{[(])}>]}",                       // false
	}

	for _, test := range tests {
		fmt.Printf("Input: %s → Valid: %v\n", test, isValidBracketString(test))
	}
}
