from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

L = 71
C = 71
bytes = set()


def find_exit(bytes):
    visited = set()
    q = [((0, 0), 0)]
    while len(q) > 0:
        top, l = q[0]
        if top == (L - 1, C - 1):
            return l

        q = q[1:]
        if top in visited:
            continue
        visited.add(top)

        ns = [
            (top[0] + 1, top[1]),
            (top[0] - 1, top[1]),
            (top[0], top[1] - 1),
            (top[0], top[1] + 1),
        ]

        for n in ns:
            if 0 <= n[0] < L and 0 <= n[1] < C and n not in bytes:
                q.append((n, l + 1))

    return -1


t = 0
for line in lines:
    parts = line.split(",")
    x = int(parts[0])
    y = int(parts[1])
    bytes.add((x, y))
    t += 1

    print(t)
    # part 2
    if t > 1024:
        l = find_exit(bytes)
        if l == -1:
            print(line)
            break

visited = set()
q = [((0, 0), 0)]
while True:
    top, l = q[0]
    if top == (L - 1, C - 1):
        print(l)
        break

    q = q[1:]
    if top in visited:
        continue
    visited.add(top)

    ns = [
        (top[0] + 1, top[1]),
        (top[0] - 1, top[1]),
        (top[0], top[1] - 1),
        (top[0], top[1] + 1),
    ]

    for n in ns:
        if 0 <= n[0] < L and 0 <= n[1] < C and n not in bytes:
            q.append((n, l + 1))
