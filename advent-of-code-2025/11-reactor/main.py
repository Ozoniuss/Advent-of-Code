from collections import Counter, defaultdict

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

graph = defaultdict(list)
for line in lines:
    parts = line.split(":")
    parts[1] = parts[1][1:]
    nparts = parts[1].split(" ")
    graph[parts[0]].extend(nparts)

# print(graph)


def dfs(graph, path, current):
    if current == "dac":
        return 1
        # print("path", path)
        if "fft" in path and "dac" in path:
            return 1
        else:
            return 0
    s = 0
    # print("ns", graph[current], "cur", current, "path", path)
    for n in graph[current]:
        if n in path:
            continue
        path.add(n)
        s += dfs(graph, path, n)
        path.remove(n)
    return s

def bfs(graph, start, end):
    s = 0
    q = [(start, 1)]
    while len(q) != 0:
        # print(q)
        nextq =[]
        nextd = defaultdict(int)
        for t in q:
            current, ways = t
            if current == end:
                s += ways
            for n in graph[current]:
                nextd[n] += ways
        
        for n, v in nextd.items():
            nextq.append((n,v))
        q = nextq
    return s

print(bfs(graph, "svr", "fft") * bfs(graph, "fft", "dac") * bfs(graph, "dac", "out"))
print(bfs(graph, "svr", "fft") * bfs(graph, "dac", "fft") * bfs(graph, "fft", "out"))

# part 1
# def dfs(graph, path, current):
#     if current == "out":
#         print("path", path)
#         return 1
#     s = 0
#     print("ns", graph[current], "cur", current, "path", path)
#     for n in graph[current]:
#         if n in path:
#             continue
#         path.append(n)
#         s += dfs(graph, path, n)
#         path.pop()
#     return s

# print(dfs(graph, ["you"], "you"))