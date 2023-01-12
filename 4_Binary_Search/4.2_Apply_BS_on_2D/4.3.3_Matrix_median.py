class Solution:
    def countSmallerThanEqualToMid (self, arr, target):
        l, r = 0, len(arr)-1
        while l<=r:
            mid = l+((r-l)//2)
            if arr[mid] <= target:
                l = mid+1
            else:
                r = mid-1
        return l

    # Take a number range, and keep searching on how many numbers exists
    # in the matrix before and after that mid.
    def median(self, matrix, R, C):
        l, r = 1, 10**9
        n, m = len(matrix), len(matrix[0])
        while l<=r:
            mid = l+((r-l)//2)
            
            numsFound = 0
            for i in range(n):
                numsFound += self.countSmallerThanEqualToMid(matrix[i], mid)

            if numsFound <= ((n*m)//2):
                l = mid+1
            else:
                r = mid-1
        return l