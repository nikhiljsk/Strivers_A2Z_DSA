# Python - N2 Solution - TLE
def inversionCount(self, arr, n):
    c = 0
    for i in range(n):
        for j in range(i+1, n):
            if arr[j] < arr[i]:
                c+=1
    return c

# Approach 2
# Merge sort
class Solution:
    def merge(self, arr, temp_arr, left, mid, right):
        i = left
        j = mid + 1
        k = left
        inv_count = 0
        while i <= mid and j <= right:
            if arr[i] <= arr[j]:
                temp_arr[k] = arr[i]
                i += 1
            else:
                temp_arr[k] = arr[j]
                inv_count += (mid-i + 1)
                j += 1
            k += 1
        while i <= mid:
            temp_arr[k] = arr[i]
            k += 1
            i += 1
        while j <= right:
            temp_arr[k] = arr[j]
            k += 1
            j += 1
        for loop_var in range(left, right + 1):
            arr[loop_var] = temp_arr[loop_var]
        return inv_count
    
    
    def mergeSort(self, arr, temp_arr, left, right):
        inv_count = 0
        if left < right:
            mid = (left + right)//2
            inv_count += self.mergeSort(arr, temp_arr, left, mid)
            inv_count += self.mergeSort(arr, temp_arr, mid + 1, right)
            inv_count += self.merge(arr, temp_arr, left, mid, right)
        return inv_count
    
    def inversionCount(self, arr, n):
        temp_arr = [0]*n
        return self.mergeSort(arr, temp_arr, 0, n-1)


arr = [1, 20, 6, 4, 5]
n = len(arr)
ob = Solution()
result = ob.inversionCount(arr, n)
print("Number of inversions are", result)
print(arr)