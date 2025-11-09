package problem2169

/* You are given two non-negative integers num1 and num2.
In one operation, if num1 >= num2, you must subtract num2 from num1, otherwise subtract num1 from num2.
For example, if num1 = 5 and num2 = 4, subtract num2 from num1, thus obtaining num1 = 1 and num2 = 4. 
However, if num1 = 4 and num2 = 5, after one operation, num1 = 4 and num2 = 1.
Return the number of operations required to make either num1 = 0 or num2 = 0. */

func CountOperations(num1, num2 int) int {
    larger, smaller := num1, num2
    operations := 0
    for larger > 0 && smaller > 0 {
        if larger < smaller {
            larger, smaller = smaller, larger
        }
        operations += larger / smaller
        larger = larger % smaller
    }
    return operations
}
