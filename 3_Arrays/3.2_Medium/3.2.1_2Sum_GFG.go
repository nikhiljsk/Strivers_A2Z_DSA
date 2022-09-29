// Python
def allPairs(self, A, B, N, M, X):
    found = dict()
    l = list()
    A = sorted(A)
    for i in range(len(B)):
        found[B[i]] = i
    for i in range(len(A)):
        if (X-A[i]) in found.keys():
            l.append([A[i], B[found[X-A[i]]]])
    return l

// Golang
func allPairs(A, B []int, N, M, X int) [][]int {
	found := make(map[int]int)
	var res [][]int
	sort.Ints(A[:])
	for i := 0; i < len(B); i++ {
		found[B[i]] = i
	}
	for i := 0; i < len(A); i++ {
		if _, ok := found[X-A[i]]; ok {
			res = append(res, []int{A[i], B[found[X-A[i]]]})
		}
	}
	return res
}