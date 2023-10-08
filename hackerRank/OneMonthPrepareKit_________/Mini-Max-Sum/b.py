#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'miniMaxSum' function below.
#
# The function accepts INTEGER_ARRAY arr as parameter.
#


def miniMaxSum(arr):
    # Write your code here
    arr.sort()
    minSum = sum(arr[:4])
    maxSum = sum(arr[1:])
    print(minSum, maxSum)


def miniMaxSum1(arr):
    # Write your code here
    minSum = sum(arr)
    maxSum = 0
    for i in range(len(arr)):
        # tong cac phan tu khac phan tu thu i va tong cac phan tu
        # vi du: arr = [1,2,3,4,5]
        # vi tri i = 0 thi sum1 = 2+3+4+5 = 14
        # vi tri i = 1 thi sum1 = 1+3+4+5 = 13
        # vi tri i = 2 thi sum1 = 1+2+4+5 = 12
        # vi tri i = 3 thi sum1 = 1+2+3+5 = 11
        # vi tri i = 4 thi sum1 = 1+2+3+4 = 10
        sum1 = sum(arr[:i] + arr[i+1:])  # arr[:i]: <i, arr[i+1:]: >i, bo i
        if sum1 < minSum:
            minSum = sum1
        if sum1 > maxSum:
            maxSum = sum1
    print(minSum, maxSum)


if __name__ == '__main__':

    arr = list(map(int, input().rstrip().split()))

    miniMaxSum(arr)
