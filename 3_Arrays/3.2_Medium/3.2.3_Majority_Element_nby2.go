// Aproach 1
// NlogN
func majorityElement(nums []int) int {
	sort.Ints(nums[:])
	return nums[len(nums)/2]
}

// Approach 2
// N
// Moore's Voting Algorithm
func majorityElement(nums []int) int {
	element, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			element = v
		}
		if v == element {
			count++
		} else {
			count--
		}
	}
	return element
}