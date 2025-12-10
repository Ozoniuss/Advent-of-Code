from collections import Counter, defaultdict

from email.policy import default
import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")



def distance(P1, P2):
    x1,y1,z1 = P1
    x2,y2,z2 = P2
    return (x2-x1) ** 2 + (y2-y1) ** 2 + (z2-z1) ** 2


distances = []
circuits = {}
currentCircuit = 0
for i in range(len(lines)-1):
    for j in range(i+1,len(lines)):
        parts1 = lines[i].split(",")
        P1 = (int(parts1[0]), int(parts1[1]), int(parts1[2]))
        parts2 = lines[j].split(",")
        P2 = (int(parts2[0]), int(parts2[1]), int(parts2[2]))
        d = distance(P1, P2)
        distances.append((d, i, j))

distances.sort(key= lambda x: x[0])
# print(distances)


res = None
for el in distances:
    _, P1, P2 = el
    if P1 not in circuits and P2 not in circuits:
        circuits[P1] = currentCircuit
        circuits[P2] = currentCircuit
        currentCircuit += 1
    elif P1 in circuits and P2 not in circuits:
        circuits[P2] = circuits[P1]
        if len(set(circuits.values())) == 1 and len(circuits.values()) == len(lines):
            print("aha", P1, P2, lines[P1], lines[P2])
            res = int(lines[P1].split(",")[0])*int(lines[P2].split(",")[0])
            break
    elif P1 not in circuits and P2 in circuits:
        circuits[P1] = circuits[P2]
        if len(set(circuits.values())) == 1 and len(circuits.values()) == len(lines):
            print("aha", P1, P2, lines[P1], lines[P2])
            res = int(lines[P1].split(",")[0])*int(lines[P2].split(",")[0])
            break
    else:
        if circuits[P1] != circuits[P2]:
            print('not ok', circuits, P1, P2)
            orig = circuits[P2]
            new = circuits[P1]
            for c in circuits:
                if circuits[c] == orig:
                    circuits[c] = new
            if len(set(circuits.values())) == 1 and len(circuits.values()) == len(lines):
                print("aha", P1, P2, lines[P1], lines[P2])
                res = int(lines[P1].split(",")[0])*int(lines[P2].split(",")[0])
                break
print(circuits, currentCircuit)
print(res)

# cnt = defaultdict(int)
# for c, v in circuits.items():
#     cnt[v] += 1
# print(cnt)

# v = list(cnt.values())
# v.sort(reverse=True)
# print(v[0]*v[1]*v[2])

exit(0)
# part 1
cap = 1000
for el in distances:
    _, P1, P2 = el
    # print(P1, P2)

    if cap == 0:
        if P1 not in circuits:
            circuits[P1] = currentCircuit
            currentCircuit += 1
        if P2 not in circuits:
            circuits[P2] = currentCircuit
            currentCircuit += 1
        continue

    if P1 not in circuits and P2 not in circuits:
        circuits[P1] = currentCircuit
        circuits[P2] = currentCircuit
        currentCircuit += 1
    elif P1 in circuits and P2 not in circuits:
        circuits[P2] = circuits[P1]
    elif P1 not in circuits and P2 in circuits:
        circuits[P1] = circuits[P2]
    else:
        if circuits[P1] == circuits[P2]:
            print('ok')
        else:
            print('not ok', circuits, P1, P2)
            orig = circuits[P2]
            new = circuits[P1]
            for c in circuits:
                if circuits[c] == orig:
                    circuits[c] = new


    cap -= 1
# print(circuits, currentCircuit)

cnt = defaultdict(int)
for c, v in circuits.items():
    cnt[v] += 1
print(cnt)

v = list(cnt.values())
v.sort(reverse=True)
print(v[0]*v[1]*v[2])