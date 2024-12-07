from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")


def solve(left, right, current):
    print(left, right, current)
    if current > left:
        return 0

    if len(right) == 0 and current == left:
        return 1

    if len(right) == 0 and current != left:
        return 0

    return solve(left, right[1:], current * right[0]) + solve(
        left, right[1:], current + right[0]
    )


def solve_2(left, right, current):
    print(left, right, current)
    if current > left:
        return 0

    if len(right) == 0 and current == left:
        return 1

    if len(right) == 0 and current != left:
        return 0

    return (
        solve_2(left, right[1:], current * right[0])
        + solve_2(left, right[1:], current + right[0])
        + solve_2(left, right[1:], int(str(current) + str(right[0])))
    )


t = 0
for line in lines:
    parts = line.split(":")
    left = int(parts[0])
    right = list(map(int, parts[1].strip().split(" ")))

    # if solve(left, right[1:], right[0]) > 0:
    # t += left
    if solve_2(left, right[1:], right[0]) > 0:
        t += left

print(t)
