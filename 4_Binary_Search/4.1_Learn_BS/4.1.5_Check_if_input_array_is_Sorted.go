// Python
def arraySortedOrNot(self, arr, n):
        prev = arr[0]
        for i in range(1, n):
            if prev > arr[i]:
                return False
            prev = arr[i]
        return True

// Golang
func arraySortedOrNot(arr []int, n int) bool{
    prev := arr[0]
    for i:=1; i<n; i++{
        if prev > arr[i]{
            return false
        }
        prev = arr[i]
    }
    return true
}

// Recursive Golang
func isSorted(arr []int) bool {
	n := len(arr)
	if n==0 || n==1 {
		return true
	}
	return arr[0] > arr[1] && isSorted(arr[1:])
}