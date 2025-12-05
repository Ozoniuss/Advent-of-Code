from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

ranges = []
ingredients = []


blocknr = 1
for line in lines:
    if line == "":
        blocknr += 1
        continue

    if blocknr == 1:
        parts = line.split("-")
        ranges.append((int(parts[0]), int(parts[1])))

    if blocknr == 2:
        ingredients.append(int(line))

cnt = 0
for i in ingredients:
    for r in ranges:
        if r[0] <= i and r[1] >= i:
            cnt += 1
            break

print(cnt)

cnt = 0
ranges.sort(key= lambda x: x[0])

current_range = (0,0)
for r in ranges:
    print(current_range, r, cnt)
    (left, right) = r
    if left > current_range[1]:
        current_range = r
        cnt += r[1] - r[0] + 1
        continue
    
    # current[0] <= left <= current[1]
    if r[1] > current_range[1]:
        cnt += r[1] - current_range[1] 
        current_range = (current_range[0], r[1])
print(current_range)

print(cnt)