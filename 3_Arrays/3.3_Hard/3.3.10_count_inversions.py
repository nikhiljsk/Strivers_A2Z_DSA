# Python - N2 Solution - TLE
def inversionCount(self, arr, n):
    c = 0
    for i in range(n):
        for j in range(i+1, n):
            if arr[j] < arr[i]:
                c+=1
    return c