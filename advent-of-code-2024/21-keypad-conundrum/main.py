from collections import Counter, defaultdict
import sys
import os
import math

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

board1 = [
"xxxxx",
"x789x",
"x456x",
"x123x",
"xx0Ax",
"xxxxx",
]
x1,y1 = 4,3

board2 = [
"xxxxx",
"xx^Ax",
"x<v>x",
"xxxxx",
]
x2,y2 = 1,3


def getNeighbours(board, x, y):
    ns = []
    for n in [
        (x-1, y, "^"),
        (x+1, y, "v"),
        (x, y+1, ">"),
        (x, y-1, "<"),
    ]:
        (i,j,dir)=n
        if board[i][j] != 'x':
            ns.append((i, j, dir))

    return ns

# bfs, works on any board
def find_shortest_paths(boardno: str, path: str) -> list[str]:
    assert boardno in ["1", "2"]
    board = None
    if boardno == "1":
        (x, y) = (x1, y1)
        q = [(x,y,"")]
        board = board1
    elif boardno == "2":
        (x, y) = (x2, y2)
        q = [(x,y,"")]
        board = board2
    pathLoc = 0

    while True:
        assert(len(q) > 0)
        nextSequence = dict()
        nextq = dict()
        diff = 0
        for n in q:
            currentLoc = pathLoc
            (x, y, currentPath) = n
            while currentLoc < len(path) and  board[x][y] == path[currentLoc]:
                currentLoc += 1

            if currentLoc == pathLoc:
                continue

            diff = currentLoc - pathLoc
            toAdd = (x,y,currentPath+"A"*diff)
            nextq[toAdd] = None
            
        if len(nextq) != 0:
            q = list(nextq)
            pathLoc += diff
            if pathLoc == len(path):
                ret = []
                for it in nextq:
                    ret.append(it[2])
                return ret

        for n in q:
            (x, y, currentPath) = n
            ns = getNeighbours(board, x, y)
            for nn in ns:
                (i, j, d) = nn
                toAdd = (i, j, currentPath + d)
                nextSequence[toAdd] = None

        q = list(nextSequence)

print(find_shortest_paths("1", "029A"))

shortestPathsBoard1 = {}
shortestPathsBoard2 = {}

for c1 in "0123456789A":
    for c2 in "0123456789A":
        shortest_paths = dict()
        for p in find_shortest_paths("1", c1+c2):
            shortest_paths[p.split("A")[1] + "A"] = None
        shortestPathsBoard1[c1+c2] = list(shortest_paths)

for c1 in "<>^vA":
    for c2 in "<>^vA":
        shortest_paths = dict()
        for p in find_shortest_paths("2", c1+c2):
            shortest_paths[p.split("A")[1] + "A"] = None
        shortestPathsBoard2[c1+c2] = list(shortest_paths)

print("b1",shortestPathsBoard1)
print("b2", shortestPathsBoard2)

for k in shortestPathsBoard2:
    if len(shortestPathsBoard2[k]) != 1:
        print(k, shortestPathsBoard2[k])

print("----")
# <A ['>^>A', '>>^A']
# >^ ['^<A', '<^A'] -- doesn't matter because both have the same solution length
# ^> ['v>A', '>vA'] -- doesn't matter which one
# vA ['^>A', '>^A'] -- because of 2, doesn't matter
# A< ['v<<A', '<v<A']
# Av ['v<A', '<vA'] -- doesn't mater
# print("ahoy", find_shortest_paths("2",">^>A"))
# print(find_shortest_paths("2",">>^A"))
# print(find_shortest_paths("2","v<<A"))
# print(find_shortest_paths("2","<v<A"))

# board2 = [
# "xxxxx",
# "xx^Ax",
# "x<v>x",
# "xxxxx",
# ]

shortestPathsBoard2["<A"] = ['>>^A']
shortestPathsBoard2[">^"] = ['<^A']
shortestPathsBoard2["^>"] = ['v>A']
shortestPathsBoard2["vA"] = ['^>A']
shortestPathsBoard2["A<"] = ['v<<A']
shortestPathsBoard2["Av"] = ['<vA']

for k in shortestPathsBoard2:
    assert(len(shortestPathsBoard2[k]) == 1)
    print(k, shortestPathsBoard2[k])


def find_shortest_paths_b2(path: str):
    sol = ""
    path = "A" + path
    for i in range(len(path)-1):
        sector = path[i:i+2]
        sol += shortestPathsBoard2[sector][0]
    
    return [sol]

print("test", find_shortest_paths_b2("<vA^>^A"))
print("test", find_shortest_paths_b2("<vA"))
print("test", find_shortest_paths_b2("^>^A"))

whatItProduces = []

# f(..A...A...A) = f(..A) + f(..A) + f(..A)
def find_shortest_paths_b2_len_times(path: str, n: int):
    cache = defaultdict(int)

    # initialize cache based off current path
    cur = ""
    for i in range(len(path)):
        if path[i] != "A":
            cur += path[i]
        else:
            cache[cur + "A"] += 1
            cur = ""

    # print("init", cache)
    for i in range(n):
        newcache = defaultdict(int)
        for k in cache:
            nextpath = find_shortest_paths_b2(k)[0]
            # print(k, nextpath)
            cur = ""
            for i in range(len(nextpath)):
                if nextpath[i] != "A":
                    cur += nextpath[i]
                else:
                    newcache[cur + "A"] += cache[k]
                    if newcache[cur + "A"] == 0:
                        newcache[cur + "A"] = 1
                    cur = ""
        cache = newcache
    s = 0
    for k, v in cache.items():
        s += len(k) * v
    return s

print("cREEP",find_shortest_paths_b2_len_times("AAA",13))

inputData = [
    "279A",
    "286A",
    "508A",
    "463A",
    "246A",
]

inputData2 = [
    "029A",
    "980A",
    "179A",
    "456A",
    "379A",
]

s = 0
total = 0
for d in inputData:
    path = d
    noBoards = 25

    sols = find_shortest_paths("1", path)
    minv = math.inf
    for sol in sols:
        minv = min(minv,find_shortest_paths_b2_len_times(sol, noBoards))
    print("d", d, minv)

    total += int(path[:-1]) * (minv)

    
print(total)