package problem3623

/* You are given a 2D integer array points, where points[i] = [xi, yi] represents 
the coordinates of the ith point on the Cartesian plane.
A horizontal trapezoid is a convex quadrilateral with at least one pair of horizontal sides 
(i.e. parallel to the x-axis). Two lines are parallel if and only if they have the same slope.
Return the number of unique horizontal trapezoids that can be formed by choosing any four distinct points from points.
Since the answer may be very large, return it modulo 10^9 + 7. */

func CountTrapezoids(points [][]int) int {
    trapezoidCount := 0
    const mod = 1_000_000_007
    
    pointsByHeight := map[int]int{}
    for _, point := range points {
        pointsByHeight[point[1]]++
    }
    
    totalPreviousPairs := 0
    for _, pointsAtHeight := range pointsByHeight {
        pairsAtHeight := pointsAtHeight * (pointsAtHeight - 1) / 2
        trapezoidCount += totalPreviousPairs * pairsAtHeight
        totalPreviousPairs += pairsAtHeight
    }
    
    return trapezoidCount % mod
}
