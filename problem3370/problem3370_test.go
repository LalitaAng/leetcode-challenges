package problem3370

import "testing"

func TestSmallestNumber(t *testing.T) {
    tests := []struct {
        n        int
        expected int
    }{
        {1, 1},      
        {2, 3},      
        {3, 3},      
        {4, 7},      
        {5, 7},      
        {7, 7},      
        {8, 15},     
        {15, 15},    
        {16, 31},    
        {31, 31},    
        {32, 63},    
        {100, 127},  
        {255, 255},  
        {256, 511},  
    }

    for _, tt := range tests {
        got := SmallestNumber(tt.n)
        if got != tt.expected {
            t.Errorf("SmallestNumber(%d) = %d; want %d", tt.n, got, tt.expected)
        }
    }
}
