from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

L = len(lines)
C = len(lines[0])

start = set()
for i in range(L):
    for j in range(C):
        if lines[i][j] == "0":
            start.add((i, j))


# visited = set()


def explore(current, already_reached):
    if lines[current[0]][current[1]] == ".":
        return 0
    val = int(lines[current[0]][current[1]])
    if val == 9:
        if current in already_reached:
            return 0
        # print("o yuea", current)
        already_reached.add(current)
        return 1

    n = []
    if current[0] > 0:
        n.append((current[0] - 1, current[1]))
    if current[0] < L - 1:
        n.append((current[0] + 1, current[1]))
    if current[1] > 0:
        n.append((current[0], current[1] - 1))
    if current[1] < C - 1:
        n.append((current[0], current[1] + 1))

    o = 0
    for neigh in n:
        if lines[neigh[0]][neigh[1]] == ".":
            continue
        if int(lines[neigh[0]][neigh[1]]) == val + 1:
            o += explore(neigh, already_reached)

    return o


# This is actually my original approach. I debugged it a lot, not understanding
# what the fuck I am overcounting. Apparently, I was doing part b solution.
def explore_original(current):
    if lines[current[0]][current[1]] == ".":
        return 0
    val = int(lines[current[0]][current[1]])
    if val == 9:
        return 1

    n = []
    if current[0] > 0:
        n.append((current[0] - 1, current[1]))
    if current[0] < L - 1:
        n.append((current[0] + 1, current[1]))
    if current[1] > 0:
        n.append((current[0], current[1] - 1))
    if current[1] < C - 1:
        n.append((current[0], current[1] + 1))

    o = 0
    for neigh in n:
        if lines[neigh[0]][neigh[1]] == ".":
            continue
        if int(lines[neigh[0]][neigh[1]]) == val + 1:
            o += explore_original(neigh)

    return o


print(start)
t = 0
for s in start:
    # score = explore(s, set())
    score = explore_original(s)
    print("s", score, s)
    t += score
print(t)
