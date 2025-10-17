package problem191

import "testing"

func TestHammingWeight(t *testing.T) {
    tests := []struct {
        name string
        input int
        want int
    }{
        {
            name:  "Example 1",
            input: 11, // binary: 1011
            want:  3,
        },
        {
            name:  "Example 2",
            input: 128, // binary: 10000000
            want:  1,
        },
        {
            name:  "Example 3",
            input: 4294967293, // binary: 11111111111111111111111111111101
            want:  31,
        },
        {
            name:  "Zero",
            input: 0,
            want:  0,
        },
        {
            name:  "All bits set",
            input: 4294967295, // all 32 bits are 1
            want:  32,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := HammingWeight(int(tt.input))
            if got != tt.want {
                t.Errorf("HammingWeight(%v) = %v; want %v", tt.input, got, tt.want)
            }
        })
    }
}
