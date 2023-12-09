import sys
import os

solvepart = None
if not sys.argv[1] == "--no-ask":
    solvepart = input("which part to solve> ")

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")


def part1():
    ...


def part2():
    ...


if solvepart == None:
    part1()
    part2()
    exit(0)

ans = part1() if solvepart == "1" else part2()
print(ans)
