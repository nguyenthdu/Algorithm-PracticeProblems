#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'reverseArray' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts INTEGER_ARRAY a as parameter.
#


def reverseArray(a):
    # Reverse Array using the slicing
    # need creating more requires more space. the exhausts more memory
    newList = a[::-1]
    # return newList
    # reverse list using two-pointer approach
    left = 0
    right = len(a)-1
    while left < right:
        # swap
        temp = arr[left]
        arr[left] = arr[right]
        arr[right] = temp
        left += 1
        right -= 1

    # using comphrehesion
    newList = [a[len(a)-i] for i in range(1, len(a)+1)]


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    arr_count = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    res = reverseArray(arr)

    fptr.write(' '.join(map(str, res)))
    fptr.write('\n')

    fptr.close()
