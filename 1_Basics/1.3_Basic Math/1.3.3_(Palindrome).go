func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	res, copyOfX := 0, x
	for x > 0 {
		res = (res * 10) + x%10
		x /= 10
	}
	fmt.Println(res)
	return copyOfX == res
}

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%v", x)
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}