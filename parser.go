package portainerparser

type Pairs map[rune]rune

var pairs = Pairs{
	'(': ')',
	'{': '}',
	'[': ']',
}

func Parse(input string) bool {
	if (len(input) % 2) > 0 {
		return false
	}

	reverse := make(Pairs, len(pairs)*2)
	for k, v := range pairs {
		reverse[v] = k
	}

	var stack []rune
	for _, c := range input {
		// Check if character is an opener.
		if _, ok := pairs[c]; ok {
			stack = append(stack, c)
			continue
		}

		if _, ok := reverse[c]; ok {
			if len(stack) == 0 {
				// Starting with closer is an error.
				return false
			}

			if stack[len(stack)-1] != reverse[c] {
				return false
			}

			// Remove from stack on success.
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
