from collections import Counter, defaultdict

from email.policy import default
import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

L = len(lines)
C = len(lines[0])

start = None
end = None

board = [list(l) for l in lines]
for i in range(L):
    for j in range(C):
        if board[i][j] == 'S':
            start = (i,j)
        if board[i][j] == 'E':
            end = (i, j)

print("start", start, "end", end)

road = 0
stk = [start]
visited = set()
visited.add(start)

m = []
ml = 999999999999
while len(stk) != 0:
    top = stk[-1]
    if top == end:
        if len(stk) < ml:
            ml = len(stk)
            m = stk.copy()
            stk.pop()
        continue
    dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]
    ns = [(top[0]+x, top[1]+y) for (x,y) in dirs]

    shouldpop = True
    for n in ns:
        if 0<=n[0]<L and 0<=n[1]<C and board[n[0]][n[1]] != '#' and n not in visited:
            stk.append(n)
            visited.add(n)
            shouldpop = False
            continue
    if shouldpop:
        stk.pop()
print(m, len(m))

# c = defaultdict(int)
# c2 = 0
# dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]
# for (tracklength, step) in enumerate(m):
#     if step == end:
#         continue
#     x,y = step
#     for (dx, dy) in dirs:
#         nx = x+dx
#         ny = y+dy
#         # not cheating
#         if board[nx][ny] != "#":
#             continue
#         for (ddx, ddy) in dirs:
#             # print(dx, dy, ddx, ddy)
#             # print(step, (nx, ny), end=" ")
#             nnx = nx + ddx
#             nny = ny + ddy
#             # print((nnx, nny))
#             if nnx < 0 or nnx >= L or nny < 0 or nny >= C:
#                 # print("broken", nx, ny)
#                 continue
#             if board[nnx][nny] == '#':
#                 # print("broken hashtag", nx, ny)
#                 continue
#             try:
#                 posinstep = m.index((nnx,nny))
#                 if posinstep > tracklength + 2:
#                     saves = posinstep - tracklength -2
#                     print("cheat saves",saves,step, m[posinstep], m.index(step), posinstep)
#                     c[saves] += 1
#                     if saves >= 100:
#                         c2 += 1
#             except:
#                 assert(0 == 1)

# calculate dp over hashtags
saves_100 = 0
def calculate_how_shorter(how_shorter, hashtag, m, ending_points):
    global saves_100
    dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]
    ns = []

    current, h = hashtag
    l, tracklength = h
    for d in dirs:
        ns.append((current[0] + d[0], current[1] + d[1]))

    l = l+1
    
    for n in ns:
        try:
            # if n in ending_points:
                # continue
            posinstep = m.index(n)
            saves = posinstep - tracklength - l
            if saves > 0:
                # print("cheat saves",saves,step, m[posinstep], m.index(step), posinstep)
                # how_shorter[saves] += 1
                # if saves >= 100:
                    # saves_100 += 1
                if n not in ending_points:
                    ending_points[n] = saves
                # if n in ending_points and ending_points[n] < saves:
                    # ending_points[n] = saves
                # else:
                    # ending_points[n] += 1
        except:
            # assert(0 == 1)
            pass

how_shorter = defaultdict(int)
dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]

# Original version from when I did not understand how this worked.
# for (tracklength, step) in enumerate(m):
#     print(step, tracklength)
#     if step == end:
#         continue

#     q = [(step, 0)]

#     # all reachable hashtags and number of steps it takes to reach them
#     reachable = {}
#     ending_points = set()

#     # build reachable
#     while(len(q)) != 0:
#         current, l = q[0]
#         x, y = current[0], current[1]
#         q = q[1:]

#         # cheat too long, regardless it finished
#         if l > 21:
#             continue

#         if not (0 <= current[0] < L and 0 <= current[1] < C):
#             continue

#         # if l == 21:
#         #     try:
#         #         posinstep = m.index(n)
#         #         saves = posinstep - tracklength - l
#         #         if saves > 0:
#         #             pass
#         #             # print("cheat saves",saves,step, m[posinstep], m.index(step), posinstep)
#         #             # how_shorter[saves] += 1
#         #             # if saves >= 100:
#         #                 # saves_100 += 1
#         #             # if n not in ending_points:
#         #                 # ending_points[n] = saves
#         #             # if n in ending_points and ending_points[n] < saves:
#         #                 # ending_points[n] = saves
#         #             # else:
#         #                 # ending_points[n] += 1
#         #     except:
#         #         # assert(0 == 1)
#         #         pass

#         for (dx, dy) in dirs:
#             nx = x+dx
#             ny = y+dy
#             if not (0 <= nx < L and 0 <= ny <C):
#                 continue
#             # if board[nx][ny] != '#':
#             #     continue
#             if (nx, ny) in reachable:
#                 continue
#             # assert(board[nx][ny] == '#')
#             q.append(((nx, ny), l+1))
#             reachable[(nx, ny)] = (l+1, tracklength)
#     ending_points = {}
#     for r in reachable:
#         calculate_how_shorter(how_shorter, (r, reachable[r]), m, ending_points)
#     for ep in ending_points:
#         how_shorter[ending_points[ep]] += 1

# vals = [i for i in how_shorter.items()]
# vals.sort(key=lambda x: x[0], reverse=True)
# print(vals)
# print(saves_100)

# Fast version that was pretty obvious once I understood.
how_much_it_saves = defaultdict(int)
over_100 = 0
for (l, step) in enumerate(m):
    print(l, step)
    for (l2, end) in enumerate(m):
        # reaches a point ahead or in the same point
        if l2 <= l:
            continue
        distance = abs(step[0] - end[0]) + abs(step[1] - end[1])
        if distance > 20:
            continue
        saves = l2 - l - distance
        if saves > 0:
            how_much_it_saves[saves] += 1
            if saves >= 100:
                over_100 += 1
# print(how_much_it_saves)
# vals = [i for i in how_much_it_saves.items()]
# vals.sort(key=lambda x: x[0], reverse=True)
# print(vals)
print(over_100)