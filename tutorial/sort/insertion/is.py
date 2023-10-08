def Insertion(list):
    for i in range(len(list)):
        key = list[i]
        j = i-1
        while j >= 0 and key < list[j]:
            list[j+1] = list[j]
            j -= 1

        list[j+1] = key


list = [64, 34, 25, 12, 22, 11, 90]
Insertion(list)
for i in list:
    print(i)
