package problem717

import "testing"

func TestIsOneBitCharacter(t *testing.T) {
    tests := []struct {
        bits []int
        want bool
    }{
        {[]int{1, 0, 0}, true},
        {[]int{1, 1, 1, 0}, false},
        {[]int{0}, true},
        {[]int{1, 0}, false},
        {[]int{0, 0}, true},
    }
    
    for _, tt := range tests {
        if got := IsOneBitCharacter(tt.bits); got != tt.want {
            t.Errorf("IsOneBitCharacter(%v) = %v, want %v", 
                tt.bits, got, tt.want)
        }
    }
}
