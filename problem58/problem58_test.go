package problem58

import "testing"

func TestLengthOfLastWord(t *testing.T) {
    tests := []struct {
        name string
        s    string
        want int
    }{
        {
            name: "Normal case",
            s:    "Hello World",
            want: 5,
        },
        {
            name: "Trailing spaces",
            s:    "Hello World   ",
            want: 5,
        },
        {
            name: "Single word",
            s:    "Lilly",
            want: 5,
        },
        {
            name: "Multiple spaces between words",
            s:    "   fly me   to   the moon  ",
            want: 4,
        },
        {
            name: "One character",
            s:    "a",
            want: 1,
        },
        {
            name: "Empty string",
            s:    "",
            want: 0,
        },
        {
            name: "Spaces only",
            s:    "     ",
            want: 0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := LengthOfLastWord(tt.s)
            if got != tt.want {
                t.Errorf("LengthOfLastWord(%q) = %d, want %d", tt.s, got, tt.want)
            }
        })
    }
}
