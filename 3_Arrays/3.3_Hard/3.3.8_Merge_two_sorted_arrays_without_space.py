import math

# O(N2)
def merge2(self,arr1,arr2,n,m):
    for i in range(n):
        if arr1[i] > arr2[0]:
            arr1[i], arr2[0] = arr2[0], arr1[i]
        
        first = arr2[0]
        j = 1
        while j<m and first > arr2[j]:
            arr2[j-1] = arr2[j]
            j+=1
        arr2[j-1] = first


# Python - Gap Method
# O(Log2(m+n) * (m+n))
def merge(self,arr1,arr2,n,m):
    gap = math.ceil(n+m/2)
    while gap > 0:
        i, j = 0, gap
        while j < n+m:
            if j<n and arr1[i]>arr1[j]:                         # If both pointers in same arr1
                arr1[i], arr1[j] = arr1[j], arr1[i]
            elif j>=n and i<n and arr1[i]>arr2[j-n]:            # If one in arr1 other in arr2
                arr1[i], arr2[j-n] = arr2[j-n], arr1[i]
            elif j>=n and i>=n and arr2[i-n]>arr2[j-n]:         # If both pointers in same arr2
                arr2[j-n], arr2[i-n] = arr2[i-n], arr2[j-n]
            i+=1
            j+=1
        if gap == 1:
            break
        gap = math.ceil(gap/2)