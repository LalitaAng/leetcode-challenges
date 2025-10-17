package problem191

// Given a positive integer n, write a function that returns the number of set bits in its binary representation (also known as the Hamming weight).

func HammingWeight(n int) int {
    count := 0
    for i := 0; i < 32; i++ {
        if n&(1<<i) != 0 {
            count++
        }
    }
    return count
}
