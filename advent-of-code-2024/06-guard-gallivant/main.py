from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

UP = (-1, 0)
DOWN = (1, 0)
LEFT = (0, -1)
RIGHT = (0, 1)


def turn_right(dir):
    if dir == UP:
        return RIGHT
    if dir == RIGHT:
        return DOWN
    if dir == DOWN:
        return LEFT
    if dir == LEFT:
        return UP


L = len(lines)
C = len(lines[0])

start = None
for i in range(L):
    for j in range(C):
        if lines[i][j] == "^":
            start = (i, j)


visited = set()
print(start)
current = start
visited.add(current)

dir = UP

total = 0
visited = set()
for i in range(L):
    for j in range(C):
        print(i, j)
        changed = False
        orig = ""
        if lines[i][j] == ".":
            orig = lines[i]
            lines[i] = lines[i][:j] + "#" + lines[i][j + 1 :]
            changed = True

        if changed:
            visited = set()
            current = start
            visited.add((start, UP))
            dir = UP
            while 0 <= current[0] < L and 0 <= current[1] < C:
                next = (current[0] + dir[0], current[1] + dir[1])
                if not (0 <= next[0] < C and 0 <= next[1] < C):
                    current = next
                    break
                if lines[next[0]][next[1]] == "#":
                    dir = turn_right(dir)
                else:
                    current = next
                    if (next, dir) in visited:
                        total += 1
                        break
                    visited.add((current, dir))
            lines[i] = orig
            changed = False
print(total)
