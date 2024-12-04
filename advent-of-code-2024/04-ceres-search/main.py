from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

print(len(lines), len(lines[0]))

t = 0
for i in range(len(lines)):
    for j in range(len(lines[0]) - 3):
        word = lines[i][j : j + 4]
        if word == "XMAS" or word == "SAMX":
            t += 1

for i in range(len(lines[0]) - 3):
    for j in range(len(lines)):
        word = lines[i][j] + lines[i + 1][j] + lines[i + 2][j] + lines[i + 3][j]
        if word == "XMAS" or word == "SAMX":
            t += 1

for i in range(len(lines[0]) - 3):
    for j in range(len(lines[0]) - 3):
        word1 = (
            lines[i][j]
            + lines[i + 1][j + 1]
            + lines[i + 2][j + 2]
            + lines[i + 3][j + 3]
        )
        word2 = (
            lines[i + 3][j]
            + lines[i + 2][j + 1]
            + lines[i + 1][j + 2]
            + lines[i][j + 3]
        )
        if word1 == "XMAS" or word1 == "SAMX":
            t += 1
        if word2 == "XMAS" or word2 == "SAMX":
            t += 1

t = 0
for i in range(len(lines) - 2):
    for j in range(len(lines[0]) - 2):
        quad = [lines[i][j : j + 3], lines[i + 1][j : j + 3], lines[i + 2][j : j + 3]]
        print(quad)
        leftu = quad[0][0]
        rightu = quad[0][2]
        leftd = quad[2][0]
        rightd = quad[2][2]
        mid = quad[1][1]
        if mid != "A":
            continue
        if leftu == rightd or leftd == rightu:
            continue
        total = leftu + rightu + leftd + rightd
        if "M" not in total or "S" not in total:
            continue
        c = dict(Counter(total))
        print(c)

        if c["M"] != 2 or c["S"] != 2:
            continue
        t += 1


print(t)
