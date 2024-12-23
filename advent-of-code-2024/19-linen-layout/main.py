from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

towels = []
designs = []
for line in lines:
    if line == "":
        continue
    elif "," in line:
        towels = line.split(", ")
    else:
        designs.append(line)

# print(towels, designs)


cache = {}


def can_make(towels, design):
    global cache
    if design in cache:
        return cache[design]
    tot = 0
    for t in towels:
        # print(t, design, design[-len(t) :])
        if t == design:
            tot += 1
        if len(t) > len(design):
            continue
        if len(t) == len(design) and t != design:
            continue
        if len(t) < len(design) and design[-len(t) :] == t:
            tot += can_make(towels, design[: -len(t)])
            # t += 1
    cache[design] = tot
    return tot


c = 0
for d in designs:
    print(d)
    c += can_make(towels, d)
print(c)
