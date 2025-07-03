package questions

func IsValid(s string) bool {
	stack := make([]rune, 0)
	dict := make(map[rune]rune)
	dict['('] = ')'
	dict['{'] = '}'
	dict['['] = ']'
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if dict[top] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
