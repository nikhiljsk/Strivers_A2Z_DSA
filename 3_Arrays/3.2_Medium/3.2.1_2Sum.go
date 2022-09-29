func twoSum(nums []int, target int) []int {
	find := make(map[int]int)
	for i, v := range nums {
		if j, ok := find[v]; ok && i != j {
			return []int{j, i}
		}
		find[target-v] = i
	}
	return []int{}
}