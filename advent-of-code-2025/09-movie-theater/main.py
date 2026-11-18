from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")


maxr = 0
t = 0
for i in range(len(lines)-1):
    for j in range(i,len(lines)):
        parts1 = lines[i].split(",")
        parts2 = lines[j].split(",")
        x1, y1 = int(parts1[0]), int(parts1[1])
        x2, y2 = int(parts2[0]), int(parts2[1])

        minx = min(x1, x2)
        miny = min(y1,y2)
        maxx = max(x1, x2)
        maxy = max(y1,y2)

        maxr = max(maxr, (abs(maxx-minx) + 1) * (abs(maxy-miny)+1))
print(maxr)

xlines = []
ylines = []
edges = {}

for i in range(len(lines)-1):
    parts1 = lines[i].split(",")
    x, y = int(parts1[0]), int(parts1[1])
    xlines.append(x)
    ylines.append(y)

horizontals = []
verticals = []
for i in range(len(lines)-1):
    j = i+1
    parts1 = lines[i].split(",")
    parts2 = lines[j].split(",")
    x1, y1 = int(parts1[0]), int(parts1[1])
    x2, y2 = int(parts2[0]), int(parts2[1])

    toadd = [(x1,y1), (x2, y2)]

    if y1 == y2:
        toadd.sort(key = lambda x: x[0])
        horizontals.append((toadd[0], toadd[1]))
    else:
        toadd.sort(key= lambda x: x[1])
        verticals.append((toadd[0], toadd[1]))

xlines = list(set(xlines))
xlines.sort()
ylines = list(set(ylines))
ylines.sort()

print(len(xlines), len(ylines),xlines, ylines)

horizontals.sort(key=lambda x: x[0][1])
verticals.sort(key=lambda x: x[1][0])
print(horizontals)
print(verticals)

visited = {}
# find a rectange that is inside the polygon
q = [(0, 1)]
visited[(0,1)] = None
while len(q) > 0:
    # print("q",q)
    cur = q[0]
    q = q[1:]

    x, y = cur[0], cur[1]
    rectx, recty = xlines[cur[0]], ylines[cur[1]]
    print(rectx, recty)
    for n in [(x+1, y), (x-1, y), (x, y+1), (x, y-1)]:
        nx, ny = n
        if 0 <= nx < len(xlines)-1 and 0<=ny<len(ylines)-1 and n not in visited:
            # only add neighbours that are still inside the polygon
            q.append(n)
            visited[n] = None
