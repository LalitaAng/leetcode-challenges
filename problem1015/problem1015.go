package problem1015

/* Given a positive integer k, you need to find the length of the smallest positive integer n such that 
n is divisible by k, and n only contains the digit 1.
Return the length of n. If there is no such n, return -1.
Note: n may not fit in a 64-bit signed integer. */

func SmallestRepunitDivByK(k int) int {
    if k % 2 == 0 || k % 5 == 0 {
        return -1
    }

    rem := 0

    for length := 1; length <= k; length++ {
        rem = (rem * 10 + 1) % k
        if rem == 0 {
            return length
        }
    }

    return -1
}
