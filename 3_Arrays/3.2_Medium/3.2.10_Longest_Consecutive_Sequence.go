func longestConsecutive(nums []int) int {
	found := make(map[int]bool, len(nums))
	for _, v := range nums {
		found[v] = true
	}

	curr, curr_len, max_len := 0, 0, 0
	for k, _ := range found {
		// Allow next for loop, only if not part of sequences found so far
		// Also don't allow any number which has a previous number
		if _, ok := found[k-1]; !ok {
			curr = k
			curr_len = 1

			// Calculate the longest sequence possible with curr
			ok_exists := found[curr+1]
			for ok_exists {
				curr += 1
				curr_len += 1
				_, ok_exists = found[curr+1]
			}

			if curr_len > max_len {
				max_len = curr_len
			}
		}
	}

	return max_len
}