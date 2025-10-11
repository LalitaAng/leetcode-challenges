package problem200

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "Single island in small grid",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "Two separate islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '1', '1'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 2,
		},
		{
			name: "All water",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "All land",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gridCopy := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]byte, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}

			result := NumIslands(gridCopy)
			if result != tt.expected {
				t.Errorf("NumIslands() = %d; want %d", result, tt.expected)
			}
		})
	}
}
