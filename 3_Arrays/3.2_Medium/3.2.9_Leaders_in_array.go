// Python solution
def leaders(self, nums, N):
    maxi, i = nums[N-1], N-1
    res = list()
    while i>=0:
        if maxi <= nums[i]:
            res.append(nums[i])
            maxi = nums[i]
        i-=1
    return res[-1::-1]