from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

print(lines)


pos = 50
c = 0
for l in lines:
    dir = l[0]
    n = int(l[1:])
    assert (dir in ["L", "R"])
    if dir == "L":
        pos = (pos - n % 100) % 100
    else:
        pos = (pos + n % 100) % 100
    if pos == 0:
        c += 1

print(c)

pos = 50
c = 0
for l in lines:
    dir = l[0]
    n = int(l[1:])
    assert (dir in ["L", "R"])
    last = pos
    if dir == "L":
        pos = (pos - n)
    else:
        pos = (pos + n)

    if pos >= 100:
        c += pos // 100
    elif pos <= 0:
        c += abs((pos-1) // 100)
        if last == 0:
            c -= 1

    pos = pos % 100

print(c)
print("===")


print(-100//100)