package problem350

/* Given two integer arrays nums1 and nums2, return an array of their intersection. 
Each element in the result must appear as many times as it shows in both arrays and you may 
return the result in any order. */

func Intersect(nums1 []int, nums2 []int) []int {
    found := make(map[int]int)
	var result []int
	for _, num := range nums1 {
		found[num]++
	}
	for _, num := range nums2 {
		if found[num] > 0 {
			result = append(result, num)
			found[num]--
		}
	}
	return result
}
