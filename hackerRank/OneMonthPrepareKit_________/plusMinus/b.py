#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'plusMinus' function below.
#
# The function accepts INTEGER_ARRAY arr as parameter.
#


def plusMinus(arr):
    # Write your code here
    pos = 0.0
    nag = 0.0
    zero = 0.0
    n = len(arr)
    for i in range(n):
        if arr[i] == 0:
            zero += 1
        elif arr[i] > 0:
            pos += 1
        else:
            nag += 1
    print(pos/n)
    print(nag/n)
    print(zero/n)


if __name__ == '__main__':
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)
