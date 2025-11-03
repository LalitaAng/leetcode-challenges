package problem1578

/* Alice has n balloons arranged on a rope. You are given a 0-indexed string colors where colors[i] is the color of the ith balloon.
Alice wants the rope to be colorful. She does not want two consecutive balloons to be of the same color, so she asks Bob for help. 
Bob can remove some balloons from the rope to make it colorful. You are given a 0-indexed integer array neededTime where 
neededTime[i] is the time (in seconds) that Bob needs to remove the ith balloon from the rope.
Return the minimum time Bob needs to make the rope colorful. */

func MinCost(colors string, neededTime []int) int {
    var ans int64 = 0
    var blockSum int64 = 0
    var blockMax int64 = 0
    n := len(colors)
    
    for i := 0; i < n; i++ {
        if i > 0 && colors[i] != colors[i-1] {
            ans += blockSum - blockMax
            blockSum = 0
            blockMax = 0
        }
        blockSum += int64(neededTime[i])
        if int64(neededTime[i]) > blockMax {
            blockMax = int64(neededTime[i])
        }
    }
    ans += blockSum - blockMax
    return int(ans)
}
