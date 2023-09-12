def LSearch(arr, n, k):
    flag = 0
    i = 0
    while i < n-1:
        if arr[i] == k:
            print(f'found at position {i}')
            flag = 1
            break
        i += 1
    if flag == 0:
        print("not found")


arr = [1, 23, 4, 5, 7]
n = 5
k = 4
LSearch(arr, n, k)
