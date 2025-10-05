package problem567

import "testing"

func TestCheckInclusion(t *testing.T) {
    tests := []struct {
        s1     string
        s2     string
        expect bool
    }{
        {"ab", "eidbaooo", true},   
        {"ab", "eidboaoo", false},  
        {"adc", "dcda", true},      

        {"a", "a", true},           
        {"a", "b", false},          
        {"abc", "bbbca", true},    
        {"abc", "ccccbbbbaaaa", false}, 
        {"abc", "cbaebabacd", true}, 
        {"", "anything", true},     
        {"longstring", "short", false}, 
    }

    for _, tt := range tests {
        got := CheckInclusion(tt.s1, tt.s2)
        if got != tt.expect {
            t.Errorf("CheckInclusion(%q, %q) = %v; want %v",
                tt.s1, tt.s2, got, tt.expect)
        }
    }
}
