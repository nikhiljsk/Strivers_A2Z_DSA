// Approach 1
// Brute force
// O(N2)
func pMaxSubArray(nums []int, target int) {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == target {
				for k := i; k <= j; k++ {
					fmt.Print(nums[k], " ")
				}
				fmt.Println()
			}
		}
	}
}

// Approach 2
// Sliding Window
// O(N2)
func pMaxSubArray2(nums []int, target int) {
	l, r := 0, -1
	sum := 0
	for l < len(nums) {
		for (r+1 < len(nums)) && (sum+nums[r+1] <= target) {
			sum += nums[r+1]
			r++
		}
		if sum == target {
			for i := l; i <= r; i++ {
				fmt.Print(nums[i], " ")
			}
			fmt.Println()
		}
		sum -= nums[l]
		l++
	}
}


// GFG Question
def pairWithMaxSum(self, arr, N):
	sum = 0
	for i in range(N-1):
		sum = max(sum, arr[i]+arr[i+1])
	return sum