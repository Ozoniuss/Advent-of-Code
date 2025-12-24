from cProfile import label
from collections import Counter

import sys
import os
import networkx as nx
import matplotlib.pyplot as plt

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")


wires = {}
wg = nx.DiGraph()

first = True
operations = []
for line in lines:
    if line == "":
        first = False
        continue
    if first:
        parts = line.split(": ")
        wires[parts[0]] = int(parts[1])
        wg.add_node(parts[0])
    else:
        operations.append(line)
        parts = line.split(" ")
        wire1 = parts[0]
        wire2 = parts[2]
        op = parts[1]
        out = parts[4]
        wg.add_edge(wire1, out, label='c')
        wg.add_edge(wire2, out, label='x')

print(operations, wires)
nx.draw(wg, with_labels=True)

for layer, nodes in enumerate(nx.topological_generations(wg)):
    # `multipartite_layout` expects the layer as a node attribute, so add the
    # numeric layer value as a node attribute
    for node in nodes:
        wg.nodes[node]["layer"] = layer

# Compute the multipartite_layout using the "layer" node attribute
pos = nx.multipartite_layout(wg, subset_key="layer", scale=5)

fig, ax = plt.subplots(figsize=(100,75))
nx.draw_networkx(wg, pos=pos, ax=ax, node_size=300,font_size=8,arrowsize=2)
ax.set_title("DAG layout in topological order")
# fig.tight_layout()
plt.show()

plt.savefig("graph.png")

next_operations = []
while len(operations) > 0:
    print(len(operations))
    next_operations = []
    for line in operations:
        parts = line.split(" ")

        wire1 = parts[0]
        wire2 = parts[2]
        op = parts[1]
        out = parts[4]

        try:
            if op == "AND":
                val = wires[wire1] & wires[wire2]
                wires[out] = val
            if op == "OR":
                val = wires[wire1] | wires[wire2]
                wires[out] = val
            if op == "XOR":
                val = wires[wire1] ^ wires[wire2]
                wires[out] = val
        except KeyError:
            next_operations.append(line)
    operations = next_operations

print(wires)
zs = []
for wire in wires.keys():
    if wire[0] == "z":
        zs.append(wire)

zs.sort()
o = ""
for z in zs:
    o = str(wires[z]) + o
print(o, int(o, 2))
