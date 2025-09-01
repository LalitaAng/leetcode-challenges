package problem20

/* Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.
	Every close bracket has a corresponding open bracket of the same type. */

func IsValid(s string) bool {
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{}

	for _, char := range s {
		if open, isClose := bracketMap[char]; isClose {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}