func isAnagram(s string, t string) bool {
	find := make([]int, 26)
	for _, v := range s {
		find[v-'a'] += 1
	}
	for _, v := range t {
		find[v-'a'] -= 1
	}
	for _, v := range find {
		if v != 0 {
			return false
		}
	}
	return true
}