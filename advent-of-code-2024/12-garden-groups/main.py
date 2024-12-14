from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

L = len(lines)
C = len(lines[0])
regions = []
starting_pos = []
for i in range(L):
    for j in range(C):
        starting_pos.append((i, j))

print(starting_pos)

visited = set()
areas = []
def explore(sp, visited: set, areas: list[list[tuple[int, int]]]):
    
    if sp in visited:
        return
    area = []
    q = [sp]
    val = lines[sp[0]][sp[1]]
    while len(q) > 0:    
        curr = q.pop()
        visited.add(curr)
        area.append(curr)
        nbs = [(curr[0]-1, curr[1]), (curr[0]+1, curr[1]), (curr[0], curr[1]-1), (curr[0], curr[1]+1)]
        for n in nbs:
            x, y = n
            print(x, y)
            if (0 <= x < L) and (0 <= y < C) and ((x, y) not in visited) and lines[x][y] == val and (x, y) not in q:
                q.append(n)
    areas.append(area)


def calculate_perimeter(area):
    t = 0 
    for curr in area:
        nbs = [(curr[0]-1, curr[1]), (curr[0]+1, curr[1]), (curr[0], curr[1]-1), (curr[0], curr[1]+1)]
        for n in nbs:
            if n not in area:
                t += 1
    return t


def consecutive_pairs_row(row):
    t = 0
    for idx in range(len(row)):
        if idx == 0:
            t += 1
            continue
        if row[idx][1] - row[idx-1][1] > 1:
            t += 1
    return t
def consecutive_pairs_col(col):
    t = 0
    for idx in range(len(col)):
        if idx == 0:
            t += 1
            continue
        if col[idx][0] - col[idx-1][0] > 1:
            t += 1
    return t
    

def calculate_sides(area):
    t = 0
    for i in range(L):
        up = []
        down = []
        for j in range(C):
            if (i, j) in area:
                if (i-1,j) not in area:
                    up.append((i-1, j))
                if (i+1, j) not in area:
                    down.append((i+1, j))
        
        t += consecutive_pairs_row(up)
        t += consecutive_pairs_row(down)
       
    for j in range(C):
        left = []
        right = []
        for i in range(L):
            if (i, j) in area:
                if (i, j-1) not in area:
                    left.append((i, j-1))
                if (i, j+1) not in area:
                    right.append((i, j+1))
        
        t += consecutive_pairs_col(left)    
        t += consecutive_pairs_col(right)

    return t    

visited = set()
areas = []
for sp in starting_pos:
    explore(sp, visited, areas)

# print(sp, visited, areas)
t = 0
for a in areas:
    # print(a, len(a), calculate_perimeter(a), calculate_sides(a))
    # print(a, len(a), calculate_perimeter(a))
    t += len(a) * calculate_sides(a)
print(t)
