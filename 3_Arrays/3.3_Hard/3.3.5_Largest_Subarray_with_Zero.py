# Python Brute force or Naive
def maxLen2(self, n, arr):
    res, s = 0, 0
    for i in range(0, len(arr)):
        s = 0
        for j in range(i, len(arr)):
            s += arr[j]
            if s == 0:
                res = max(j-i+1, res)
    return res


def maxLen(n, arr):
        prefix_sum = dict()
        s = 0
        res = 0
        for i in range(0, len(arr)):
            s += arr[i]
            if s == 0:
                res = i+1
            else:
                if s in prefix_sum.keys():
                    res = max(res, i-(prefix_sum[s]))
                else:
                    prefix_sum[s] = i
        return res
    
arr = [15, -2, 2, -8, 1, 7, 10, 23]
print(maxLen(len(arr), arr))