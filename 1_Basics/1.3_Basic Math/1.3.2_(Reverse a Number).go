func reverse(x int) int {
	res := 0
	intMax := 2147483647
	intMin := -2147483648
	for x != 0 {
		r := x % 10
		x = x / 10
		if (res > intMax/10) || (res == intMax/10 && r > 7) {
			return 0
		}
		if (res < intMin/10) || (res == intMin/10 && r < -8) {
			return 0
		}
		res = (res * 10) + r
	}
	return res
}