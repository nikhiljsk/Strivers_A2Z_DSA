// Python code
def findKRotation(self,nums,  n):
	l, r = 0, len(nums)-1
	last = nums[r]
	while l <= r:
		mid = l + (r-l)//2
		if nums[mid] > last:
			l = mid + 1
		else:
			r = mid - 1

	return l