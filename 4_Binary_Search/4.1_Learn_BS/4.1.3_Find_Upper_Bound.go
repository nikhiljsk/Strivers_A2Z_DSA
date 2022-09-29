// Python
def findFC(A,N,X):
    l, r = 0, N-1
    while l<=r:
        mid = l+((r-l)//2)
        if A[mid]>X:
            r = mid-1
        elif A[mid]<X:
            l = mid+1
        else:
            return A[mid], A[mid]
    if l==0:
        return -1, A[l]
    if l >= N:
        return A[l-1], -1
    return A[l-1], A[l]

def getFloorAndCeil(arr, n, x):
    arr = sorted(arr)
    return findFC(arr, n, x)

// Golang
func findFC(A []int, N, X int) (int, int) {
	l, r := 0, N-1
	for l <= r {
		mid := l + ((r - l) / 2)
		if A[mid] > X {
			r = mid - 1
		} else if A[mid] < X {
			l = mid + 1
		} else {
			return A[mid], A[mid]
		}
	}
	if l == 0 {
		return -1, A[l]
	} else if l >= N {
		return A[l-1], -1
	}
	return A[l-1], A[l]
}

func getFloorAndCeil(arr []int, N, X int) (int, int) {
	sort.Ints(arr[:])
	return findFC(arr, N, X)
}