from collections import Counter

import sys
import os

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


def find_shortest_paths(board: list[str], start: tuple[int, int], path: str) -> list[str]:
    # print(board, start, path)
    (x, y) = start
    q = [(x,y,"")]
    pathLoc = 0

    while True:
        # print(length, q)
        assert(len(q) > 0)
        nextSequence = set()
        nextq = set()
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
            nextq.add(toAdd)
            # if toAdd not in nextq:
                # nextq.append(toAdd)
            
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
                if toAdd not in nextSequence:
                    nextSequence.add(toAdd)

        q = list(nextSequence)

print(find_shortest_paths(board1, (x1,y1), "029A"))


shortestPathsBoard1 = {}
shortestPathsBoard2 = {}

for c1 in "0123456789A":
    for c2 in "0123456789A":
        shortest_paths = set()
        for p in find_shortest_paths(board1, (x1,y1), c1+c2):
            shortest_paths.add(p.split("A")[1] + "A")
        shortestPathsBoard1[c1+c2] = list(shortest_paths)

for c1 in "<>^vA":
    for c2 in "<>^vA":
        shortest_paths = set()
        for p in find_shortest_paths(board2, (x2,y2), c1+c2):
            shortest_paths.add(p.split("A")[1] + "A")
        shortestPathsBoard2[c1+c2] = list(shortest_paths)


print(shortestPathsBoard1["A0"])
print(shortestPathsBoard1["02"])
print(shortestPathsBoard1["29"])
print(shortestPathsBoard1["9A"])
print(shortestPathsBoard1)
print(shortestPathsBoard2)

def find_shortest_paths_len(boardKind, path: str) -> int:

    useBoard = None
    if boardKind == "1":
        useBoard = shortestPathsBoard1
    else:
        useBoard = shortestPathsBoard2 
    
    l = 0
    path = "A" + path
    for i in range(len(path)-1):
        sector = path[i:i+2]
        l += (min(list(map(len, useBoard[sector])))) 
    return l

shortestMathMinimized = {}
for c1 in "<>^vA":
    for c2 in "<>^vA":
        allHere = {}
        shortenedPaths = []
        # minl = 99999999999
        for p in shortestPathsBoard2[c1+c2]:
            ll = find_shortest_paths_len("2", p)
            allHere[p] = ll
            # if ll < minl:
            #     minl = ll
        for p in allHere:
            # if allHere[p] == minl:
            shortenedPaths.append(p)
        # pick 1 because the 4 that have 2 left are
# >^ ['<^A', '^<A']
# ^> ['>vA', 'v>A']
# vA ['^>A', '>^A']
# Av ['v<A', '<vA'] so it doesnt really matter
        shortestMathMinimized[c1+c2] = shortenedPaths
    
print("initial",shortestPathsBoard2)
print("filtered", shortestMathMinimized)
for k in shortestMathMinimized:
    if len(shortestMathMinimized[k]) == 2:
        print(k, shortestMathMinimized[k])
# exit(0)

print(find_shortest_paths_len("1", "029A"))

def find_shortest_paths_with_len(boardKind, path: str):

    useBoard = None
    if boardKind == "1":
        useBoard = shortestPathsBoard1
    else:
        useBoard = shortestMathMinimized 

    solutions = [""]

    path = "A" + path
    for i in range(len(path)-1):
        sector = path[i:i+2]
        newSols = []
        for s in solutions:
            for b in useBoard[sector]:
                newSols.append(s + b)
        solutions = newSols
    
    return solutions

def find_shortest_paths_with_len_v2(board, path: str):
    sol = ""
    path = "A" + path
    for i in range(len(path)-1):
        sector = path[i:i+2]
        sol += shortestMathMinimized[sector][0]
    
    return [sol]

print("tt", find_shortest_paths_with_len("1", "029A"))
start = "<^"
for i in range(10):
    print(len(start))
    n = find_shortest_paths_with_len_v2(None, start)
    start = n[0]
start = ">^"
for i in range(10):
    print(len(start))
    n = find_shortest_paths_with_len_v2(None, start)
    start = n[0]

exit(0)

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

# s = 0
# for d in inputData2:

#     paths = [[d]]
#     boards=[(board1, (x1, y1)), (board2, (x2,y2))]
#     boards.append((board2, (x2, y2)))

#     for board in boards:
#         nextPaths = []
#         for path in paths[-1]:
#             nextPaths.extend(find_shortest_paths(board[0], board[1], path))
#         paths.append(nextPaths)
#         print("all", paths)
#     print("number of total solution", len(paths[-1]))

#     minl = 999999999999999
#     minp = None
#     for p in paths[-1]:
#         if len(p) < minl:
#             minl = len(p)
#             minp = p
    
#     print(int(d[:-1]), len(minp))
#     s += int(d[:-1]) * len(minp)
# print(s)


## testing length
# s = 0
# for d in inputData:

#     paths = [[d]]
#     boards=[(board1, (x1, y1)), (board2, (x2,y2)), (board2, (x2, y2))]

#     for board in boards:
#         nextPaths = []
#         for path in paths[-1]:
#             nextPaths.extend(find_shortest_paths(board[0], board[1], path))
#         paths.append(nextPaths)
#         # print("all", paths)
#     print("number of total solution remaining for", d, len(paths[-1]))

#     minlen = 9999999999999
#     for p in paths[-1]:
#         minlen = min(minlen, find_shortest_paths_len("2", p))
#     s += int(d[:-1]) * minlen
# print(s)


s = 0
for d in inputData2:

    paths = [[d]]
    boards=["1"]
    boards.extend(["2"]* 25)
    print("d", d)
    for idx, board in enumerate(boards):

        fn = None
        if board == "1":
            fn = find_shortest_paths_with_len
        else:
            fn = find_shortest_paths_with_len_v2

        print(idx)
        nextPaths = []
        pathLens = {}
        for path in paths[-1]:
            pathLens[path] = fn(board, path)
            # nextPaths.extend(find_shortest_paths_with_len(board, path))
            print("len", len(path))
        print("candidatres", len(pathLens))


        # print(d, list(pathLens.values()))
        minval = min(list(pathLens.values()))

        for p in pathLens:
            if pathLens[p] == minval:
                nextPaths.extend(fn(board, p))
        

        paths.append(nextPaths)
        print(len(nextPaths), nextPaths)
    # print(d, paths[-1])

    minl = 999999999999999
    minp = None
    for p in paths[-1]:
        if len(p) < minl:
            minl = len(p)
            minp = p
    
    print(int(d[:-1]), len(minp))
    s += int(d[:-1]) * len(minp)
print(s)