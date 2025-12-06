from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read()
lines = data.split("\n")
items = []
print(lines)



for l in lines:
    newline = []
    parts = l.split(" ")
    for part in parts:
        if part != "":
            newline.append(part)
    items.append(newline)
for item in items:
    if len(item) != len(items[0]):
        assert("lol")

line = lines[-1]
operands = []
cur = ""
for i in range(len(line)):
    if i+1 == len(line) and cur != "":
        operands.append(cur + " ")
        break
    if i+1 == len(line) or line[i+1] != " ":
        operands.append(cur)
        cur = ""
    else:
        cur += line[i]

# print(operands)
elements = []
for i in range(len(lines)-1):
    line = lines[i]
    newline = []
    cidx = 0
    for op in operands:
        width = len(op)
        newline.append(line[cidx:cidx+width])
        cidx += width + 1
    elements.append(newline)
    assert(len(newline) == len(operands))
# print(elements)


## part 1
gt = 0
for j in range(len(items[0])):
    op = items[-1][j]
    s = 0
    if op == "*":
        s = 1
    for i in range(len(items)-1):
        el = int(items[i][j])
        if op == "*":
            s *= el
        else:
            s += el

    gt += s
print("p1", gt)
###

## part 2
# print(elements)
gt = 0
for j in range(len(items[0])):
    op = items[-1][j]
    s = 0
    if op == "*":
        s = 1
    width = len(operands[j])
    intops = []
    for cw in range(width-1,-1,-1):
        eltobuild = ""
        for i in range(len(elements)):
            el = elements[i][j]
            el = el[cw:cw+1]
            if el != " " and el != "":
                eltobuild += el
        intops.append(int(eltobuild))
    # print("iops", intops)
    
    for el in intops:
        if op == "*":
            s *= el
        else:
            s += el  
    # for i in range(len(elements)):
    #     el = elements[i][j]
    #     print(i, el)
    #     intops = []
    #     for i in range(width):
    #         strop = ""


    gt += s
print("p2", gt)