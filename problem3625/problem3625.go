package problem3625

/* You are given a 2D integer array points where points[i] = [xi, yi] represents the coordinates of 
the ith point on the Cartesian plane.
Return the number of unique trapezoids that can be formed by choosing any four distinct points from points.
A trapezoid is a convex quadrilateral with at least one pair of parallel sides. Two lines are parallel if and only if 
they have the same slope. */

func CountTrapezoids(points [][]int) int {
    n := len(points)
    infiniteSlope := 1e9 + 7
    slopeToIntercepts := make(map[float64][]float64)
    segmentMidpointToSlopes := make(map[float64][]float64)
    result := 0
    
    for i := 0; i < n; i++ {
        x1, y1 := points[i][0], points[i][1]
        for j := i + 1; j < n; j++ {
            x2, y2 := points[j][0], points[j][1]
            deltaX := x1 - x2
            deltaY := y1 - y2
            
            var slope, intercept float64
            if x2 == x1 {
                slope = infiniteSlope
                intercept = float64(x1)
            } else {
                slope = float64(y2-y1) / float64(x2-x1)
                intercept = float64(y1*deltaX-x1*deltaY) / float64(deltaX)
            }
            
            midpointHash := float64((x1+x2)*10000 + (y1 + y2))
            slopeToIntercepts[slope] = append(slopeToIntercepts[slope], intercept)
            segmentMidpointToSlopes[midpointHash] = append(segmentMidpointToSlopes[midpointHash], slope)
        }
    }
    
    for _, intercepts := range slopeToIntercepts {
        if len(intercepts) == 1 {
            continue
        }
        frequency := make(map[float64]int)
        for _, intercept := range intercepts {
            frequency[intercept]++
        }
        totalSum := 0
        for _, count := range frequency {
            result += totalSum * count
            totalSum += count
        }
    }
    
    for _, slopes := range segmentMidpointToSlopes {
        if len(slopes) == 1 {
            continue
        }
        frequency := make(map[float64]int)
        for _, slope := range slopes {
            frequency[slope]++
        }
        totalSum := 0
        for _, count := range frequency {
            result -= totalSum * count
            totalSum += count
        }
    }
    
    return result
}
