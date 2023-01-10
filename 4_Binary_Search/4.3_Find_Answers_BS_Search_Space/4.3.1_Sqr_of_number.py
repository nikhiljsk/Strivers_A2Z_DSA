def floorSqrt(self, x):
    if x<=1:
        return 1
    l, r = 1, x
    while l<=r:
        mid = l+((r-l)//2)
        mid_square = mid*mid
        if mid_square < x:
            l = mid+1
        elif mid_square > x:
            r = mid-1
        else:
            return mid
    return r