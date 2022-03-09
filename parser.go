package portainerparser

import "fmt"

func Parse(input string) bool {
	var stack []rune
	for _, c := range input {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')':
			if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		fmt.Println(stack)
	}
	return true
}
