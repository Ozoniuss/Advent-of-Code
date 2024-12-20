from collections import Counter

from operator import add
import re
import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

board = [list(line) for line in lines]
L = len(board)
C = len(board[0])

E = None
S = None

for i in range(L):
    for j in range(C):
        if board[i][j] == 'S':
            S = (i, j)
        if board[i][j] == 'E':
            E = (i, j)

UP = (-1, 0)
DOWN = (1, 0)
LEFT = (0, -1)
RIGHT = (0, 1)

visited = {}
mini = 999999999999

def bfs_itr(visited):
    global mini
    q = [(S, RIGHT, 0)]
    while len(q) > 0:
        top = q[0]
        q = q[1:]

        ps, dr, val = top

        if ps == E and val < mini:
            mini = val
            continue

        if (ps, dr) in visited and val > visited[(ps, dr)]:
            continue
        visited[(ps, dr)] = val

        ns = [
            ((ps[0] + UP[0], ps[1] + UP[1]), UP),
            ((ps[0] + DOWN[0], ps[1] + DOWN[1]), DOWN),
            ((ps[0] + LEFT[0], ps[1] + LEFT[1]), LEFT),
            ((ps[0] + RIGHT[0], ps[1] + RIGHT[1]), RIGHT),
            ]
        for nanddir in ns:
            n, ndir = nanddir
            if board[n[0]][n[1]] == '#':
                continue

            if (dr[0] + ndir[0], dr[1] + ndir[1]) == (0,0):
                continue
            if dr == ndir:
                q.append((n, dr, val + 1))
            else:
                q.append((n, ndir, val + 1001))

# the following takes way too long, doesn't seem to be efficient in python
sys.setrecursionlimit(100000)
def explore_dfs(ps, dr, val, visited):
    print(ps, val)
    global mini
    if ps == E and val < mini:
        mini = val
        return

    if (ps, dr) in visited and val > visited[(ps, dr)]:
        return
    
    visited[(ps, dr)] = val
    ns = [
        ((ps[0] + UP[0], ps[1] + UP[1]), UP),
        ((ps[0] + DOWN[0], ps[1] + DOWN[1]), DOWN),
        ((ps[0] + LEFT[0], ps[1] + LEFT[1]), LEFT),
        ((ps[0] + RIGHT[0], ps[1] + RIGHT[1]), RIGHT),
          ]
    for nanddir in ns:
        n, ndir = nanddir
        if board[n[0]][n[1]] == '#':
            continue

        if (dr[0] + ndir[0], dr[1] + ndir[1]) == (0,0):
            continue
        if dr == ndir:
            explore(n, dr, val + 1, visited)
        else:
            explore(n, ndir, val + 1001, visited)



# dfs style approach, useful for part b because you have the stack and can
# trace the path.
# Could probably be optimized by including the list of neighbours of the current
# node to the stack but with the help of visited is probably unnecessary.
def dfs_itr(visited):
    
    # part b only
    # mval = 7036
    # mval = 11048
    mval = 108504
    tiles = {}

    global mini
    val = 0
    s = [(S, RIGHT, 0)]

    while len(s) > 0:
        ps, dr, val = s[-1]
        # print(ps,dr,val)

        if ps == E and val == mval:
            for frame in s:
                tiles[frame[0]] = frame[2]

        if (ps, dr) in visited and val > visited[(ps, dr)]:
            s.pop()
            continue

        if ps == E and val < mini:
            mini = val
            print(mini)
            s.pop()
            visited[(ps, dr)] = val
            continue

        visited[(ps, dr)] = val

        ns = [
            ((ps[0] + UP[0], ps[1] + UP[1]), UP),
            ((ps[0] + DOWN[0], ps[1] + DOWN[1]), DOWN),
            ((ps[0] + LEFT[0], ps[1] + LEFT[1]), LEFT),
            ((ps[0] + RIGHT[0], ps[1] + RIGHT[1]), RIGHT),
            ]
        
        shouldpop = True
        for nanddir in ns:
            n, ndir = nanddir
            if board[n[0]][n[1]] == '#':
                continue

            if (dr[0] + ndir[0], dr[1] + ndir[1]) == (0,0):
                continue

            if dr == ndir:
                if (n, dr) in visited and val + 1 >= visited[(n, dr)]:
                    if n in tiles and tiles[n] == val + 1:
                        for fr in s:
                            if fr not in tiles:
                                tiles[fr[0]] = fr[2]
                    continue
                s.append((n, dr, val + 1))
                shouldpop = False
                break
            else:
                if (n, ndir) in visited and val + 1001 >= visited[(n, ndir)]:
                    if n in tiles and tiles[n] == val + 1001:
                        for fr in s:
                            if fr not in tiles:
                                tiles[fr[0]] = fr[2]
                    continue
                s.append((n, ndir, val + 1001))
                shouldpop = False
                break
        if shouldpop:
            s.pop()
    print("total tiles", len(tiles))

print(S, E)
# bfs_itr(visited)
dfs_itr(visited)
print(mini)