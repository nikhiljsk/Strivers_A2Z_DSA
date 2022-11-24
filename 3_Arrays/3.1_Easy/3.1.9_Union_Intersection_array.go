// Union
def mergeArrays(self, a, b, m, n):
    i, j = 0, 0
    res = list()
    prev = None
    while i < m and j < n:
        if a[i] < b[j]:
            if prev != a[i]:
                res.append(a[i])
                prev = a[i]
            i += 1
        elif a[i] > b[j]:
            if prev != b[j]:
                res.append(b[j])
                prev = b[j]
            j += 1
        else:
            if prev != a[i]:
                res.append(a[i])
                prev = a[i]
            i += 1
            j += 1

    while i < m:
        if prev != a[i]:
            res.append(a[i])
            prev = a[i]
        i += 1

    while j < n:
        if prev != b[j]:
            res.append(b[j])
            prev = b[j]
        j += 1
    return res


// Intersection
def intesectArrays(a, b, m, n):
    i, j = 0, 0
    res = list()
    prev = None
    while i < m and j < n:
        if a[i] < b[j]:
            i += 1
        elif a[i] > b[j]:
            j += 1
        else:
            res.append(a[i])
            i += 1
            j += 1
    return res
	