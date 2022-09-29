func subarraySum(nums []int, k int) int {
	mydict := make(map[int]int)
	sum, count := 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == k {
			count++
		}
		if _, ok := mydict[sum-k]; ok {
			count += mydict[sum-k]
		}
		mydict[sum] += 1
	}

	return count
}