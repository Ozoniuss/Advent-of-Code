from collections import Counter, defaultdict

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

antennas = defaultdict(list)
L = len(lines)
C = len(lines[0])
for i in range(L):
    for j in range(C):
        if lines[i][j] != '.':
            antennas[lines[i][j]].append((i,j))
print(antennas)

antinodes = set()
for a, positions in antennas.items():
    for i in range(len(positions) - 1):
        for j in range(i+1, len(positions)):
            left = (2*positions[i][0] - positions[j][0], 2*positions[i][1] - positions[j][1])
            right = (2*positions[j][0] - positions[i][0], 2*positions[j][1] - positions[i][1])

            if 0 <=left[0] < L and 0 <= left[1] < C:
                antinodes.add(left)
            if 0 <=right[0] < L and 0 <= right[1] < C:
                antinodes.add(right)
print(len(antinodes))

antinodes = set()
for a, positions in antennas.items():
    for i in range(len(positions) - 1):
        for j in range(i+1, len(positions)):
            left = (2*positions[i][0] - positions[j][0], 2*positions[i][1] - positions[j][1])
            right = (2*positions[j][0] - positions[i][0], 2*positions[j][1] - positions[i][1])

            current = positions[i]
            other = positions[j]
            while 0 <=current[0] < L and 0 <= current[1] < C:
                antinodes.add(current)
                aux = current
                current = (2*current[0] - other[0], 2*current[1] - other[1])
                other = aux
            
            current = positions[j]
            other = positions[i]
            while 0 <=current[0] < L and 0 <= current[1] < C:
                antinodes.add(current)
                aux = current
                current = (2*current[0] - other[0], 2*current[1] - other[1])
                other = aux

print(len(antinodes))