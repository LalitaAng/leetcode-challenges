package problem2528

import "runtime/debug"

var (
    sumArray  [100001]int 
    diffArray [100001]int 
)

func MaxPower(stations []int, r, k int) int64 {
    minPower := stations[0]
    totalPower := 0
    
    for i, stationPower := range stations {
        sumArray[i] = 0
        minPower = min(minPower, stationPower)
        totalPower += stationPower
    }
    
    n := len(stations)
    sumArray[n] = 0
    
    for i := range n {
        leftBound := max(i-r, 0)    
        rightBound := min(i+r+1, n)  
        sumArray[leftBound] += stations[i]
        sumArray[rightBound] -= stations[i]
    }
    
    low := minPower        
    high := totalPower + k   
    result := 0               
    
    for low <= high {
        mid := low + (high-low)/2
        if canAchieve(n, mid, r, k) {  
            result = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    
    return int64(result)
}

func canAchieve(n, targetMin, r, k int) bool {
    copy(diffArray[:n], sumArray[:n])
    
    currentSum := 0
    remaining := k
    
    for i := range diffArray[:n] {
        currentSum += diffArray[i]
        
        if currentSum < targetMin {
            needed := targetMin - currentSum
            if remaining < needed {
                return false
            }
            remaining -= needed
            endPos := min(2*r+i+1, n)
            diffArray[endPos] -= needed
            currentSum += needed
        }
    }
    
    return true
}

func init() {
    debug.SetMemoryLimit(9_000_000)
}
    