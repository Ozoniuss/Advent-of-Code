from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")


wires = {}


first = True
operations = []
for line in lines:
    if line == "":
        first = False
        continue
    if first:
        parts = line.split(": ")
        wires[parts[0]] = int(parts[1])
    else:
        operations.append(line)
        # parts = line.split(" ")

        # wire1 = parts[0]
        # wire2 = parts[2]
        # op = parts[1]
        # out = parts[4]
        # print(wire1, wire2, op, out)

        # if op == "AND":
        #     val = wires[wire1] & wires[wire2]
        #     wires[out] = val
        # if op == "OR":
        #     val = wires[wire1] | wires[wire2]
        #     wires[out] = val
        # if op == "XOR":
        #     val = wires[wire1] ^ wires[wire2]
        #     wires[out] = val


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


zs = []
for wire in wires.keys():
    if wire[0] == "z":
        zs.append(wire)

zs.sort()
o = ""
for z in zs:
    o += str(wires[z])
print(int(o, 2))
