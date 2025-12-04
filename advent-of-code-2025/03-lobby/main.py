from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")


def findJoltage(row: str):
    lmax, rmax = -1, -1
    lpos = -1

    for l in range(len(row)-1):
        lval = int(row[l])
        if lval > lmax:
            lmax = lval
            lpos = l

    for r in range(lpos+1,len(row)):
        rval = int(row[r])
        if rval > rmax:
            rmax = rval

    
    return int(str(lmax)+str(rmax))


def findJoltage12(row: str):

    maxs = [-1] * 12
    pos = [-1] * 13
    pos[0] = 0

    for it in range(12):
        for l in range(pos[it], (len(row)-11)+it):
            # print(l)
            val = int(row[l])
            if val > maxs[it]:
                maxs[it] = val
                pos[it+1] = l+1

        print(maxs)

        
    return int("".join(list(map(str, maxs))))

s = 0
for line in lines:
    # jlt = findJoltage(line)
    jlt = findJoltage12(line)
    print(jlt)
    s += jlt

print(s)
