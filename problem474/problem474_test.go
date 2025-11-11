package problem474

import "testing"

func TestFindMaxForm(t *testing.T) {
    tests := []struct {
        strs []string
        m, n int
        want int
    }{
        {[]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
        {[]string{"10", "0", "1"}, 1, 1, 2},
        {[]string{"10"}, 1, 1, 1},
        {[]string{}, 1, 1, 0}, 
        {[]string{"0", "0", "0"}, 3, 0, 3}, 
    }
    
    for _, tt := range tests {
        got := FindMaxForm(tt.strs, tt.m, tt.n)
        if got != tt.want {
            t.Errorf("FindMaxForm(%v, %d, %d) = %d, want %d",
                tt.strs, tt.m, tt.n, got, tt.want)
        }
    }
}