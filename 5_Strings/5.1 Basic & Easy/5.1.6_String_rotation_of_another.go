// Approach 1
// N-square, N
func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

// Approch 2
// N-square, 1
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if len(s) == 0 {
		return true
	}

	n := len(s)
	for i := 0; i < n; i++ {
		found := false
		for j := i; j < i+n; j++ {
			k := j - i
			if s[j%n] != goal[k] {
				found = false
				break
			} else {
				found = true
			}
		}
		if found {
			return true
		}
	}
	return false
}