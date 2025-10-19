package problem8

/* Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.
The algorithm for myAtoi(string s) is as follows:
Whitespace: Ignore any leading whitespace (" ").
Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. 
If no digits were read, then the result is 0.
Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. 
Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.
Return the integer as the final result. */

func MyAtoi(s string) int {
	// Constants for 32-bit integer range
	const INT_MAX = 2147483647
	const INT_MIN = -2147483648

	n := len(s)
	if n == 0 {
		return 0
	}

	i := 0
	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	result := 0
	for i < n {
		ch := s[i]
		if ch < '0' || ch > '9' {
			break
		}
		digit := int(ch - '0')

		if result > INT_MAX/10 || (result == INT_MAX/10 && digit > 7) {
			if sign == 1 {
				return INT_MAX
			}
			return INT_MIN
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}
