// Tried Python solution - Failed
// Works only for +ve integers as input?
def lenOfLongSubarr (self, arr, n, k) :
        l, r, maxLen, count = 0, 0, 0, 0
        while r<n:
            if count+arr[r] <= k:
                count += arr[r]
                r+=1
            else:
                count -= arr[l]
                l+=1

            if maxLen < r-l and count==k:
                maxLen = r-l
        return maxLen

// Working Approach in Python
def lenOfLongSubarr (self, arr, n, k) :
    mydict = dict()
    sum = 0
    maxLen = 0

    for i in range(n):
        sum += arr[i]
        if (sum == k):
            maxLen = i + 1
        elif (sum - k) in mydict:
            maxLen = max(maxLen, i - mydict[sum - k])

        if sum not in mydict:
            mydict[sum] = i

    return maxLen

// Presum, similar to above but using different array
def lenOfLongSubarr (self, arr, n, k) : 
    pre = list()
    count = 0
    for i in range(len(arr)):
        count+=arr[i]
        pre.append(count)

    found = dict()
    maxLen = 0
    for i in range(len(arr)):
        if pre[i] == k:
            maxLen = i+1
        elif pre[i]-k in found.keys():
            maxLen = max(maxLen, i-found[pre[i]-k])
        if pre[i] not in found.keys():
            found[pre[i]] = i
    return maxLen


// Approach in Golang
func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

// The intuition here: For every element so far found, just accumulate the sum,
// And if the sum-k is found somewhere in the hashmap, that means,
// You could exlude that previous found sum, and see if it is greater than 
// so far found maxLen of subArrya
func lenOfLongSubarr(arr []int, n, k int) int {
	mydict := make(map[int]int)
	sum, maxLen := 0, 0

	for i := 0; i < n; i++ {
		sum += arr[i]
		if sum == k {
			maxLen = i + 1
		} else if _, ok := mydict[sum-k]; ok {
			maxLen = max(maxLen, i-mydict[sum-k])
		}

		if _, ok := mydict[sum]; !ok {
			mydict[sum] = i
		}
	}

	return maxLen
}



