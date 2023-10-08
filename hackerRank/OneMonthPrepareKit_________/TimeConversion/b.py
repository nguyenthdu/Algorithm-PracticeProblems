#!/bin/python3

from datetime import datetime
import math
import os
import random
import re
import sys

#
# Complete the 'timeConversion' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#


def timeConversion(s):
    # Write your code here
    if s[-2:] == 'AM':
        if s[:2] == '12':
            return '00' + s[2:-2]  # 12:00:00AM->00:00:00
        else:
            return s[:-2]  # 05:45:54AM->05:45:54
    else:
        if s[:2] == '12':
            return s[:-2]  # 12:00:00PM->12:00:00
        else:
            return str(int(s[:2]) + 12) + s[2:-2]  # 05:45:54PM->17:45:54


# use strptime() and strftime()


def timeConversion(s):
    return datetime.strptime(s, '%I:%M:%S%p').strftime('%H:%M:%S')
# '%I:%M:%S%p'->12-hour clock
# '%H:%M:%S'->24-hour clock


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = timeConversion(s)

    fptr.write(result + '\n')

    fptr.close()
