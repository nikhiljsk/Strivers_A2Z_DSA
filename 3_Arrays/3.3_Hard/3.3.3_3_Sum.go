// Brute force approach
// TLE 309/311
type triplet struct {
	unique [3]int
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	found := make(map[triplet]bool) // You can't have []int as keys in map, hence triplet
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					temp := [3]int{nums[i], nums[j], nums[k]}
					sort.Ints(temp[:])
					if _, ok := found[triplet{temp}]; !ok {
						res = append(res, []int{nums[i], nums[j], nums[k]})
						found[triplet{temp}] = true
					}
				}
			}
		}
	}
	return res
}

// Approach 2
// N2, 1
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
			l, h, sum := i+1, len(nums)-1, 0-nums[i]
			for l < h {
				if nums[l]+nums[h] == sum {
					res = append(res, []int{nums[i], nums[l], nums[h]})
					for (l < h) && (nums[l] == nums[l+1]) {
						l++
					}
					for (l < h) && (nums[h] == nums[h-1]) {
						h--
					}
					l++
					h--
				} else if nums[l]+nums[h] < sum {
					l++
				} else {
					h--
				}
			}
		}
	}
	return res
}
