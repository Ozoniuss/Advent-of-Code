from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

rules = {}
seq = []
appseq = False

for line in lines:

    if line == "":
        appseq = True
        continue

    if not appseq:
        parts = line.split("|")
        left = int(parts[0])
        right = int(parts[1])
        if left not in rules:
            rules[left] = [right]
        else:
            rules[left].append(right)
    else:
        seq.append(list(map(int, line.split(","))))


print(seq)
t = 0
for s in seq:
    ok = True
    # i was lucky tbh here because of the properties of the list
    for i in range(0, len(s) - 1):
        left, right = s[i], s[i + 1]
        if left not in rules or right not in rules[left]:
            ok = False
            break
    if ok:
        print(s[len(s) // 2])
        t += s[len(s) // 2]
print(t)

t = 0
for s in seq:

    c = s[:]

    hasswapped = True
    swaps = 0

    while hasswapped:
        hasswapped = False
        # assumed that if a|b and b|c then a|c would be explicit
        for i in range(0, len(c) - 1):
            left, right = c[i], c[i + 1]
            if left not in rules or right not in rules[left]:
                c[i], c[i + 1] = c[i + 1], c[i]
                hasswapped = True
                swaps += 1
                break
    print(c)
    if swaps > 0:
        t += c[len(c) // 2]
print(t)
