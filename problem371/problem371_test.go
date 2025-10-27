package problem371

import "testing"

func TestGetSum(t *testing.T) {
    tests := []struct {
        a, b   int
        expect int
    }{
        {1, 2, 3},
        {-2, 3, 1},
        {5, 7, 12},
        {0, 0, 0},
        {-5, -3, -8},
        {-10, 5, -5},
        {123, 456, 579},
    }

    for _, tt := range tests {
        result := GetSum(tt.a, tt.b)
        if result != tt.expect {
            t.Errorf("GetSum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expect)
        }
    }
}
