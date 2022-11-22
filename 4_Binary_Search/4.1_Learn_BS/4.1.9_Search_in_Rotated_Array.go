// Approach 1
func findPivot(nums []int) int {
	l, r := 0, len(nums)-1
	last := nums[r]
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > last {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func bSearch(nums []int, target int) (bool, int) {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true, mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false, l
}

func search(nums []int, target int) int {
	pivot := findPivot(nums)
	fmt.Println("Pivot", pivot)
	isFound, indx := bSearch(nums[:pivot], target)
	if isFound {
		return indx
	}
	isFound, indx = bSearch(nums[pivot:], target)
	if isFound {
		return indx + pivot
	}
	return -1
}

// Approach 2
// Tricky
public int search(int[] nums, int target) {
    if (nums == null || nums.length == 0) {
        return -1;
    }
    
    /*.*/
    int left = 0, right = nums.length - 1;
    //when we use the condition "left <= right", we do not need to determine if nums[left] == target
    //in outside of loop, because the jumping condition is left > right, we will have the determination
    //condition if(target == nums[mid]) inside of loop
    while (left <= right) {
        //left bias
        int mid = left + (right - left) / 2;
        if (target == nums[mid]) {
            return mid;
        }
        //if left part is monotonically increasing, or the pivot point is on the right part
        if (nums[left] <= nums[mid]) {
            //must use "<=" at here since we need to make sure target is in the left part,
            //then safely drop the right part
            if (nums[left] <= target && target < nums[mid]) {
                right = mid - 1;
            }
            else {
                //right bias
                left = mid + 1;
            }
        }

        //if right part is monotonically increasing, or the pivot point is on the left part
        else {
            //must use "<=" at here since we need to make sure target is in the right part,
            //then safely drop the left part
            if (nums[mid] < target && target <= nums[right]) {
                left = mid + 1;
            }
            else {
                right = mid - 1;
            }
        }
    }
    return -1;
}