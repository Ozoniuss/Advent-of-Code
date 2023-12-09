import sys

solvepart = None
if len(sys.argv) >= 2 and not sys.argv[1] == "--no-ask":
    solvepart = input("which part to solve> ")

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

history = [[int(num) for num in line.split()] for line in lines]


def findLast(numbers):
    differences = [numbers[i] - numbers[i - 1] for i in range(1, len(numbers))]
    if all(diff == 0 for diff in differences):
        return numbers[-1]
    return numbers[-1] + findLast(differences)


def findFirst(numbers):
    differences = [numbers[i] - numbers[i - 1] for i in range(1, len(numbers))]
    if all(diff == 0 for diff in differences):
        return numbers[0]
    return numbers[0] - findFirst(differences)


def part1():
    c = 0
    for hist in history:
        c += findLast(hist)
    print(c)


def part2():
    c = 0
    for hist in history:
        c += findFirst(hist)
    print(c)


if solvepart == None:
    part1()
    part2()
    exit(0)

ans = part1() if solvepart == "1" else part2()
print(ans)
