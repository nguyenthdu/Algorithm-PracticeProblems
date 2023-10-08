def IS(list):
    for i in range(0, len(list)):
        for j in range(i+1, len(list)):
            if list[i] > list[j]:
                list[i], list[j] = list[j], list[i]


list = [5, 3, 6, 1]
IS(list)
for i in list:
    print(i)
