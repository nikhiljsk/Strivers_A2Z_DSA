//URL: https://practice.geeksforgeeks.org/problems/maximize-number-of-1s0905/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=maximize-number-of-1s

// Tried this in Python - Failed!
def findZeroes(arr, n, m) :
    local, glob, count, last_found_index = 0, 0, 0, 0
    for i in range(0, n):
        if arr[i]==1:
            local+=1
            glob = max(local, glob)
        elif arr[i]==0 and count<m:
            local+=1
            glob = max(local, glob)
            count+=1
            last_found_index = i
        elif arr[i]==0 and count==m and count!=0:
            local = i-last_found_index
            count = 1
            glob = max(local, glob)
            last_found_index = i
        else:
            local=0

    return glob

// The above failed in for input
// 22
// 0 1 0 1 0 1 1 1 1 0 0 0 1 0 0 0 1 0 1 1 0 1
// 5

// Sliding window technique in Python
def findZeroes(arr, n, m) :
    zero_count, max_ones, l, r = 0, 0, 0, 0
    while r<n:
        if zero_count <= m:
            if arr[r] == 0:
                zero_count += 1
            r += 1
        if zero_count > m:
            if arr[l] == 0:
                zero_count -= 1 
            l += 1
            
        if r-l > max_ones:
            max_ones = r-l
            
    return max_ones

// Sliding window technique in Golang
func findZeros(arr []int, n, m int) int {
	zero_count, maxones, l, r := 0, 0, 0, 0
	for r < n {
		if zero_count <= m {
			if arr[r] == 0 {
				zero_count++
			}
			r++
		}
		if zero_count > m {
			if arr[l] == 0 {
				zero_count--
			}
			l++
		}

		if (r - l) > maxones {
			maxones = (r - l)
		}
	}
	return maxones
}