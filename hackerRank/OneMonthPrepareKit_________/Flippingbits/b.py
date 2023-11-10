#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'flippingBits' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts LONG_INTEGER n as parameter.
#


def flippingBits(n):
    # Write your code here
    return n ^ 0xffffffff  # 0xffffffff nghĩa là 2^32 - 1
# using <<


def flippingBits(n):
    return n ^ ((1 << 32) - 1)
    # giải thích từng bước
    # n^ nghĩa là xor duyệt từng bit của n và 1
    # 1 << 32 nghĩa là 1 dịch trái 32 bit, tức là 2^32


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input().strip())

    for q_itr in range(q):
        n = int(input().strip())

        result = flippingBits(n)

        fptr.write(str(result) + '\n')

    fptr.close()
