package problem2169

import "testing"

func TestCountOperations(t *testing.T) {
	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{"example_1", 2, 3, 3},
		{"example_2", 10, 10, 1},
		{"num1_greater", 10, 5, 2},
		{"num2_greater", 5, 10, 2},
		{"one_zero", 0, 7, 0},
		{"both_zero", 0, 0, 0},
		{"num1_much_larger", 100, 10, 10},
		{"coprime_numbers", 7, 5, 5},
		{"uneven_pair", 15, 4, 7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := CountOperations(tc.num1, tc.num2)
			if result != tc.expected {
				t.Errorf("CountOperations(%d, %d) = %d; want %d",
					tc.num1, tc.num2, result, tc.expected)
			}
		})
	}
}
