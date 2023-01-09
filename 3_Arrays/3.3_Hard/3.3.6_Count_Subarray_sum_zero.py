from collections import defaultdict

# Python Brute force or Naive
def subsetXOR2(self, arr, N, K): 
    res, s = 0, 0
    for i in range(0, len(arr)):
        s = 0
        for j in range(i, len(arr)):
            s = s^arr[j]
            if s == K:
                res += 1
    return res

def subsetXOR(arr, N, K): 
    res, xr = 0, 0
    pre_xor = defaultdict(int)
    for i in range(0, len(arr)):
        xr ^= arr[i]
        if xr == K:
            res += 1
        else:
            y = xr^K
            if y in pre_xor.keys():
                res += pre_xor[y]
            pre_xor[xr] += 1
    return res


arr = [5, 6, 7, 8, 9]
print(subsetXOR(arr, len(arr), 5))