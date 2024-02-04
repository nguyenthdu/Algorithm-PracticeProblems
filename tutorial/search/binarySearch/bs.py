def BSearch(list, key: int) -> int:
    left = 0
    right = len(list)-1
    while left <= right:
        mid = int((left+right)/2)
        if list[mid] == key:
            return mid
        elif list[mid] < key:
            left = mid + 1
        else:
            right = mid-1
    return -1


list = [1, 2, 3, 4, 5]
print(BSearch(list, 5))
