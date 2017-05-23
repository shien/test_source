#!/bin/env python

#
# - 1000
#
# - 500
#
# - 275
#
# - 1

def get_number(actual, range):
    upper = max(range)
    lower = min(range)
    mid = (upper - lower)  / 2 + lower
    if mid < actual:
        print("higher_range:" + str(mid + 1) + ":" + str(upper))
        j = mid + 1
        higher_range = []
        while j < upper:
           higher_range.append(j)
           j = j + 1
        get_number(actual, higher_range)
    elif mid > actual:
        print("lower_range:" + str(lower) + ":" + str(mid - 1))
        i = lower
        lower_range = []
        while i < mid -1:
           lower_range.append(i)
           i = i + 1
        get_number(actual, lower_range)
    else:
        print actual

if __name__ == "__main__":
    actual = 782 
    first_range = range(1, 1000)
    
    get_number(actual, first_range)
