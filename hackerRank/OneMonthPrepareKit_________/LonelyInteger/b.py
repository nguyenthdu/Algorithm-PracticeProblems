#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'lonelyinteger' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY a as parameter.
#


def lonelyinteger(a):
    # Write your code here
    a.sort()
    for i in range(0, len(a)-1, 2):
        if a[i] != a[i+1]:
            return a[i]
    return a[-1]


def lonelyinteger2(a):
    m = {}
    for i in a:
        if i in m:
            m[i] += 1
        else:
            m[i] = 1
    for i in m:
        if m[i] == 1:
            return i


def lonelyinteger3(a):
    return 2*sum(set(a))-sum(a)


def lonelyinteger4(a):
    res = 0
    for i in a:
        res ^= i
    return res


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    a = list(map(int, input().rstrip().split()))

    result = lonelyinteger(a)

    fptr.write(str(result) + '\n')

    fptr.close()
