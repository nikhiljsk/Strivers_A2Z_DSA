// Approach 1
// 3N, 2N
func frequencySort(s string) string {
	fqMap := make(map[rune]int)
	mxFre := 0
	for _, ch := range s {
		fqMap[ch] += 1
		if mxFre < fqMap[ch] {
			mxFre = fqMap[ch]
		}
	}
	revFqMap := make(map[int][]rune)
	for k, v := range fqMap {
		revFqMap[v] = append(revFqMap[v], k)
	}
	var res string
	for mxFre > 0 {
		if chars, ok := revFqMap[mxFre]; ok {
			for _, ch := range chars {
				for i := mxFre; i > 0; i-- {
					res += string(ch)
				}
			}
		}
		mxFre -= 1
	}
	return res
}

// Approach 2
// NlogN, N
func frequencySort(s string) string {
	if len(s) < 2 {
		return s
	}
	dict := [256]int{}
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	b := []byte(s) //type conversion

	sort.Slice(b, func(i, j int) bool {
		if dict[b[i]] == dict[b[j]] {
			return b[i] < b[j] //for same freq, sort based on element's value
			// < or >, doesn't matter here. Cause any element that appear equal number of times can be returned.
		}
		return dict[b[i]] > dict[b[j]] //sort based on freq
	})

	return string(b) //type conversion
}

// GFG
func frequencySort(nums []int) []int {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dict[nums[i]]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if dict[nums[i]] == dict[nums[j]] {
			return nums[i] < nums[j] //for same freq, sort based on element's value
			// < or >, doesn't matter here. Cause any element that appear equal number of times can be returned.
		}
		return dict[nums[i]] > dict[nums[j]] //sort based on freq
	})

	return nums
}