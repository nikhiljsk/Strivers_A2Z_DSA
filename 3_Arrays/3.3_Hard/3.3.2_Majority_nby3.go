// Approach 1
// N, N
func majorityElement(nums []int) []int {
	count := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := count[v]; ok {
			count[v] += 1
		} else {
			count[v] = 1
		}
	}

	majority := len(nums) / 3
	res := make([]int, 0)
	for k, v := range count {
		if v > majority {
			res = append(res, k)
		}
	}

	return res
}

// Approach 2
// Extended Moore's Voting Algo
func majorityElement(nums []int) []int {
	e1, e2, c1, c2 := 0, 0, 0, 0
	for _, v := range nums {
		if e1 == v {
			c1++
		} else if e2 == v {
			c2++
		} else if c1 == 0 {
			e1 = v
			c1 = 1
		} else if c2 == 0 {
			e2 = v
			c2 = 1
		} else {
			c1--
			c2--
		}
	}

	c1, c2 = 0, 0
	for _, v := range nums {
		if v == e1 {
			c1++
		} else if v == e2 {
			c2++
		}
	}

	res := make([]int, 0)
	if c1 > len(nums)/3 {
		res = append(res, e1)
	}
	if c2 > len(nums)/3 {
		res = append(res, e2)
	}

	return res
}