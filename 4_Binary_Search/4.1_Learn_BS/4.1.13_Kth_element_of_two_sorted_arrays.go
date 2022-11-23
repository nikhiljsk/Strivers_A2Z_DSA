// Python
// https://www.youtube.com/watch?v=nv7F4PiLUzo&list=PLgUwDviBIf0p4ozDR_kJJkONnb1wdx2Ma&index=66
// Tricky
def kthElement(self,  arr1, arr2, n, m, k):
    if n > m:
        return self.kthElement(arr2, arr1, m, n, k)
    
    l = max(0, k-m)
    r = min(n, k)
    
    while l<=r:
        cut1 = (l+r) >> 1
        cut2 = k-cut1
        
        if cut1 == 0:
            l1 = -sys.maxsize - 1
        else :
            l1 = arr1[cut1-1]
        
        if cut2 == 0:
            l2 = -sys.maxsize - 1
        else:
            l2 = arr2[cut2-1]
            
        if cut1 == n:
            r1 = sys.maxsize
        else:
            r1 = arr1[cut1]
        
        if cut2 == m:
            r2 = sys.maxsize
        else:
            r2 = arr2[cut2]
        
        if (l1<=r2) and (l2<=r1):
            return max(l1, l2)
        elif l1 > r2:
            r = cut1-1
        else:
            l = cut1+1
            
    return -1