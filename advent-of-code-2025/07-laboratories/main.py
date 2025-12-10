from collections import Counter, defaultdict

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

x = 0
y = lines[0].index('S')

s = [(x, y)]
cx = 0

cnt = 0
while len(s) > 0:
    nexts = []
    print(s)
    for pos in s:
        cx, cy = pos
        if cx > len(lines):
            loop = False
            break
        
        npos = (cx+1, cy)
        if cx + 1 >= len(lines):
            continue
        if lines[cx+1][cy] == '.' and npos not in nexts:
            nexts.append(npos)
        if lines[cx+1][cy] == '^':
            cnt += 1
            nposs =  [(cx+1, cy-1), (cx+1, cy+1)]

            for p in nposs:
                cxx, cyy = p
                if cyy < 0 or cyy >= len(lines[cx]):
                    continue
                if p not in nexts:
                    nexts.append(p)
    s = nexts
print(cnt)

timelines = {(x,y): 1}
cnt = 0
while len(timelines) > 0:
    print(timelines)
    ntimelines = defaultdict(int)
    for pos in timelines.keys():
        cx, cy = pos
        ct = timelines[pos]
        npos = (cx+1, cy)
        if cx + 1 >= len(lines):
            cnt += timelines[pos]
            continue

        if lines[cx+1][cy] == '.':
            ntimelines[npos] += ct
            # if npos not in ntimelines:
            #     ntimelines[npos] = ct
            # else:
            #     ntimelines[npos] += ct
        if lines[cx+1][cy] == '^':
            nposs =  [(cx+1, cy-1), (cx+1, cy+1)]
            for p in nposs:
                cxx, cyy = p
                if cyy < 0 or cyy >= len(lines[cx]):
                    continue
                ntimelines[p] += ct
                # if p not in ntimelines:
                #     ntimelines[p] = ct
                # else:
                #     ntimelines[p] += ct
    timelines = ntimelines
print(cnt)

