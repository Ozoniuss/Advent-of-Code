from collections import Counter, defaultdict

from email.policy import default
from logging import root
import sys
import os
from time import sleep

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

# L = 7
# C = 11
#
L = 103
C = 101

robots = []
for line in lines:
    parts = line.split(" ")

    ipos = list(map(int, parts[0].split("=")[1].split(",")))
    v = list(map(int, parts[1].split("=")[1].split(",")))

    ipos = [ipos[1], ipos[0]]
    v = [v[1], v[0]]

    robots.append((ipos, v))


def print_robots(robots, idx):
    # o = ""
    # for i in range(L):
    #     for j in range(C):
    #         c = 0
    #         for r in robots:
    #             if i == r[0][0] and j == r[0][1]:
    #                 c += 1
    #         if c == 0:
    #             o += "."
    #         else:
    #             o += str(c)
    #     o += "\n"
    # print(o)
    with open("tree.txt", "a") as f:
        o = ""
        for i in range(L):
            for j in range(C):
                c = 0
                for r in robots:
                    if i == r[0][0] and j == r[0][1]:
                        c += 1
                if c == 0:
                    o += "."
                else:
                    o += str(c)
            o += "\n"
        f.write(o)
        f.write("\n")
        f.write(str(idx))
        f.write("\n\n")


def print_robots_v2(robots, idx, visited):
    with open("tree.txt", "a") as f:
        o = ""
        for i in range(L):
            for j in range(C):
                c = visited[(i, j)]
                if c == 0:
                    o += "."
                else:
                    o += "#"
            o += "\n"
        f.write(o)
        f.write("\n")
        f.write(str(idx))
        f.write("\n\n")


def count_zeros(robots):
    t = 0
    for i in range(L):
        for j in range(C):
            c = 0
            for r in robots:
                if i == r[0][0] and j == r[0][1]:
                    c += 1
            if c == 0:
                t += 1
    return t


# print(robots, len(robots))
# print_robots(robots, 0)

min = 1000000000000000000

# use less to find tree, search for ######### or sth
for i in range(50000):
    print(i)
    visited = defaultdict(int)
    for r in robots:
        cpos, vel = r
        cpos[0] = (cpos[0] + vel[0]) % L
        cpos[1] = (cpos[1] + vel[1]) % C
        visited[(cpos[0], cpos[1])] += 1
    print_robots_v2(robots, i + 1, visited)
    # print(i)
# print_robots(robots, 0)

# print(robots)
clu, cld, cru, crd = (0, 0, 0, 0)
for robot in robots:
    cpos = robot[0]
    if cpos[0] < L // 2 and cpos[1] < C // 2:
        clu += 1
    if cpos[0] < L // 2 and cpos[1] > C // 2:
        cru += 1
    if cpos[0] > L // 2 and cpos[1] < C // 2:
        cld += 1
    if cpos[0] > L // 2 and cpos[1] > C // 2:
        crd += 1

print(clu, cld, cru, crd, clu * cld * cru * crd)


# a = [[1,2], [3,4]]
# x, y = a
# print(x, y)
# x[1] = 69
# print(a)
# print(1 % /7)
# print(-1 % 7)
