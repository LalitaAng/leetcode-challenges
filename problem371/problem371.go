package problem371

// Given two integers a and b, return the sum of the two integers without using the operators + and -.

func GetSum(a int, b int) int {
        for b != 0 {
        sum := a ^ b 
        carry := (a & b) << 1 
        a = sum
        b = carry
    }
    return a
}
