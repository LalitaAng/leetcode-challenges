package problem91

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"single valid", "2", 1},
		{"two digits valid separately", "12", 2},    
		{"contains zero valid", "226", 3},           
		{"invalid starting zero", "06", 0},
		{"has invalid zero", "100", 0},              
		{"multiple zeros", "101", 1},                 
		{"long valid", "11106", 2},                  
		{"empty string", "", 0},                      
		{"ends with zero valid", "2101", 1},         
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumDecodings(tt.s)
			if got != tt.want {
				t.Errorf("NumDecodings(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
