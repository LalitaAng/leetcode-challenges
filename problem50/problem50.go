package problem50

// Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

func MyPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n < 0 {
        x = 1 / x
        n = -n
    }
    result := 1.0
    for n > 0 {
        if n % 2 == 1 {
            result *= x
        }
        x *= x
        n /= 2
    }
    return result
}
