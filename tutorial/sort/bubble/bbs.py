def BubleSort(list):
    for i in range(len(list)-1):
        for j in range(len(list)-i-1):
            if list[j] > list[j+1]:
                list[j], list[j+1] = list[j+1], list[j]


list = [64, 34, 25, 12, 22, 11, 90]
BubleSort(list)
for i in list:
    print(i)
