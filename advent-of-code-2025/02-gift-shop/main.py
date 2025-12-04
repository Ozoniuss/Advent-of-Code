from collections import Counter

import sys
import os

import bisect

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

def getDivisor(ndigits):
    if ndigits <= 1:
        return None
    if ndigits %2 != 0:
        return None
    return 10 ** (ndigits // 2) + 1

def generateValidNumbers(top):
    valids = []
    for i in range(top):
        mul = i+1
        divisor =  10 ** (mul) + 1
        for k in range(10 ** (mul-1), divisor-1):
            valids.append(divisor * k)
    return valids

# print(generateValidNumbers(3))

idranges = lines[0].split(",")
validnrs = generateValidNumbers(5)
s = 0
for idrange in idranges:
    l = idrange.split("-")[0]
    r = idrange.split("-")[1]

    li = int(l)
    ri = int(r)

    if li > ri:
        l, r = r, l
        li, ri = ri, li

    afterfirst = bisect.bisect_right(validnrs, li-1)
    beforelast = bisect.bisect_left(validnrs, ri+1)

    valids = []
    for i in range(afterfirst, beforelast):
        valids.append(validnrs[i])
    s += sum(valids)
print(s)

s = 0
for idrange in idranges:
    l = idrange.split("-")[0]
    r = idrange.split("-")[1]

    li = int(l)
    ri = int(r)

    if li > ri:
        l, r = r, l
        li, ri = ri, li

    for n in range(li, ri+1):
        nstr = str(n)
        for ngroups in range(2, len(nstr)+1):
            if len(nstr) % ngroups != 0:
                continue
            gsize = len(nstr) // ngroups
            curr = None
            allEqual = True
            for k in range(ngroups):
                # print(gsize, k*gsize, (k+1)*gsize, ngroups, len(nstr))
                val = nstr[k*gsize:(k+1)*gsize]
                if curr == None:
                    curr = val
                else:
                    if val != curr:
                        allEqual = False
                        break
            if allEqual:
                s += n
                break
print(s)