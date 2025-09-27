package problem118

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{
			name:    "0 rows",
			numRows: 0,
			want:    [][]int{},
		},
		{
			name:    "1 row",
			numRows: 1,
			want:    [][]int{
				{1},
			},
		},
		{
			name:    "2 rows",
			numRows: 2,
			want: [][]int{
				{1},
				{1, 1},
			},
		},
		{
			name:    "5 rows",
			numRows: 5,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			name:    "6 rows",
			numRows: 6,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Generate(tt.numRows)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate(%d) = %v; want %v", tt.numRows, got, tt.want)
			}
		})
	}
}
