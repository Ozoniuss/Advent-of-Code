from collections import Counter

from re import L
import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

R = len(lines)
C = len(lines[0])
s = 0
for i in range(R):
    for j in range(C):
        if lines[i][j] != "@":
            continue
        cnt=0
        for n in [
            (i+1,j),
            (i-1,j),
            (i,j+1),
            (i,j-1),
            (i+1,j-1),
            (i-1,j+1),
            (i+1,j+1),
            (i-1,j-1),
        ]:
            (x, y) = n
            if x < 0 or x >= R or y < 0 or y >= C:
                continue
            if lines[x][y]=="@":
                cnt += 1
        
        if cnt < 4:
            s += 1
print(s)

s = 0
removed = True
while removed:
    removed = False
    for i in range(R):
        for j in range(C):
            if lines[i][j] != "@":
                continue
            cnt=0
            for n in [
                (i+1,j),
                (i-1,j),
                (i,j+1),
                (i,j-1),
                (i+1,j-1),
                (i-1,j+1),
                (i+1,j+1),
                (i-1,j-1),
            ]:
                (x, y) = n
                if x < 0 or x >= R or y < 0 or y >= C:
                    continue
                if lines[x][y]=="@":
                    cnt += 1
            
            if cnt < 4:
                lines[i] = lines[i][:j] + '.'+ lines[i][j+1:]
                s += 1
                print(lines[i])
                removed = True

print(s)
