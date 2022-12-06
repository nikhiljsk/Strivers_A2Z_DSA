func isIsomorphic(s string, t string) bool {
	sMap, tMap := map[byte]int{}, map[byte]int{}
	for i, _ := range s {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		} else {
			sMap[s[i]] = i + 1
			tMap[t[i]] = i + 1
		}
	}
	return true
}