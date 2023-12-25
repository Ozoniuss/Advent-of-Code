import networkx as nx

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

"""
Dissapointed to turn to python and networkx, but cutting algorithms seemed
pretty hard to implement and I was quite tired to wake up at 6:45am every
day for the past 25 days. I just wanted to get it done with.

I had other ideas and I still wanted to do it in Go for quite a while, but
it wasn't really worth it. Besides, Go doesn't have any good graphing libraries,
and I figured other people would just use libraries like yesterday's problem
(which is obviously not unfair) so I just went to Python.

Anyway, it was really fun and I'm glad that I was able to wake up every day to
do the problem. But, I'm quite happy it's over. I didn't achieve my goal of
being in the top leaderboard, but hopefully next year when I will do it in
Python I will at least get one global star. Go is definitely not the langauge
to speedrun this contest.
"""

g = nx.Graph()

for line in lines:
    parts = line.split(": ")
    s = parts[0]
    ns = [n for n in parts[1].split(" ")]

    for n in ns:
        g.add_edge(s, n, capacity=1.0)

print(g.number_of_edges())
print(g.number_of_nodes())

size, cutset = nx.minimum_cut(g, "snr", "qgb")
for e1 in list(g.nodes):
    for e2 in list(g.nodes):
        if e1 != e2:
            size, cutset = nx.minimum_cut(g, e1, e2)
            if size == 3:
                print(len(cutset[0]) * len(cutset[1]))
