package problem3370

/* You are given a positive number n.
Return the smallest number x greater than or equal to n, such that the binary representation of x contains only set bits */

func SmallestNumber(n int) int {
    x := 1
    for x < n {
        x = (x << 1) | 1
    }
    return x
}
