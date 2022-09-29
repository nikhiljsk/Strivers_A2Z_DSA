// Python
def findFloor(self,A,N,X):
    l, r = 0, N-1
    while l<=r:
        mid = l+((r-l)//2)
        if A[mid]>X:
            r = mid-1
        elif A[mid]<X:
            l = mid+1
        else:
            return mid
    return l

// Golang
func findFloor(A []int, N, X int) int {
	l, r := 0, N-1
	for l <= r {
		mid := l + ((r - l) / 2)
		if A[mid] > X {
			r = mid - 1
		} else if A[mid] < X {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}