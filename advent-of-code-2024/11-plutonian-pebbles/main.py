from collections import Counter, defaultdict

from re import L
import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

current = lines[0].split(" ")
for i in range(25):
    next = []
    for el in current:
        if el == "0":
            next.append("1")
        elif len(el) % 2 == 0:
            next.extend([str(int(el[: len(el) // 2])), str(int(el[len(el) // 2 :]))])
        else:
            next.append(str(2024 * int(el)))
    current = next
print(len(current))

current = lines[0].split(" ")
els = defaultdict(int)
for el in current:
    els[el] += 1
print(els)
for i in range(75):
    newd = defaultdict(int)
    print(i)
    next = []
    for el, v in els.items():
        if el == "0":
            newd["1"] += v
        elif len(el) % 2 == 0:
            newd[str(int(el[: len(el) // 2]))] += v
            newd[str(int(el[len(el) // 2 :]))] += v
        else:
            newd[str(2024 * int(el))] += v
    els = newd.copy()

t = 0
for _, v in els.items():
    t += v
print(t)
