package portainerparser

type pair struct {
	open  rune
	close rune
}

var pairs = []pair{
	{open: '(', close: ')'},
	{open: '{', close: '}'},
	{open: '[', close: ']'},
}

func Parse(input string) bool {
	var stack []rune
	for _, c := range input {
		for _, pair := range pairs {
			if c == pair.open {
				stack = append(stack, c)
			}
			if c == pair.close {
				// Check if previous opened matches this closed pair.
				if stack[len(stack)-1] != pair.open {
					return false
				}
				// Remove from stack on success.
				stack = stack[:len(stack)-1]
			}
		}
	}
	return true
}
