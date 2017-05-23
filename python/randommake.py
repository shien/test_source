#!/bin/python

import random




if "__main__" == __name__:
    rlist = []
    for i in range(0,20):
        rlist.append(random.randint(0,200))


    with open("test.txt", "w") as file:
        for rvalue, i in enumerate(rlist, 0):
            file.write(str(rvalue) + "\n")
