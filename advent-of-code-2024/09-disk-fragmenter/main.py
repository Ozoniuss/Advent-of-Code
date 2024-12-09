from collections import Counter, defaultdict

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")
# 00...111...2...333.44.5555.6666.777.888899
# 0099811188827773336446555566
# 009981118882777333644655556665
blocks = []

converted = ""
current_idx = 0
isFree = False
current_id = 0
for c in lines[0]:
    if not isFree:
        blocks.append([current_id, int(c)])
        current_idx += int(c)
        current_id += 1
        isFree = True
        continue
    if isFree:
        blocks.append([".", int(c)])
        current_idx += int(c)
        isFree = False
        continue

print(blocks)

pos_last_val = len(blocks) - 1  # input ends with non-free space
# out = ""
# for idx, pairs in enumerate(blocks):

#     # print(blocks)
#     if idx > pos_last_val:
#         break
#     if pairs[0] != "." and pairs[1] != 0:
#         for j in range(pairs[1]):
#             out += str(pairs[0])
#             out += "-"
#     else:
#         needs = pairs[1]
#         # print(needs)
#         while needs > 0:
#             # we're taking from a previous place to fill empty stuff
#             if pos_last_val <= idx:
#                 break
#             if blocks[pos_last_val][1] != 0:
#                 out += str(blocks[pos_last_val][0])
#                 out += "-"
#                 blocks[pos_last_val][1] -= 1
#                 needs -= 1
#             else:
#                 pos_last_val -= 2
#     # print(out)
#     # print(blocks)

# # print(blocks)
# # print(out)
# t = 0
# for i, v in enumerate(out.strip("-").split("-")):
#     t += int(i) * int(v)
# print(t)


# 00...111...2...333.44.5555.6666.777.888899
# 0099.111...2...333.44.5555.6666.777.8888..
# 0099.1117772...333.44.5555.6666.....8888..
# 0099.111777244.333....5555.6666.....8888..
# 00992111777.44.333....5555.6666.....8888..

out = ""
ok = True
blocks2 = [b.copy() for b in blocks]

pos_last_val = len(blocks) - 1

i = pos_last_val
while i > -1:
    print(i)
    # print("blocks[i]", blocks[i], i)
    # print("blocks[i]", blocks2[i], i)
    if blocks2[i][0] == ".":
        i -= 1
        continue
    needs = blocks2[i][1]
    for j in range(len(blocks2)):
        if blocks2[j][0] != ".":
            continue
        if j > i:
            break
        assert blocks2[j][0] == "."
        if blocks2[j][1] == needs:
            # print("blocks[i] changed", blocks2[i])
            blocks2[j][0] = blocks2[i][0]
            blocks2[i][0] = "."  # hope this would work
            break
        if blocks2[j][1] > needs:
            # print("blocks[i] added", blocks2[i])
            blocks2[j][0] = blocks2[i][0]
            new_empty = blocks2[j][1] - needs
            blocks2[j][1] = needs
            blocks2[i][0] = "."  # hope this would work
            blocks2 = blocks2[: j + 1] + [[".", new_empty]] + blocks2[j + 1 :]
            i += 1
            break
    i -= 1
    # print(blocks2)

cidx = 0
total = 0
for block in blocks2:
    for j in range(block[1]):
        if block[0] != ".":
            total += cidx * int(block[0])
        cidx += 1
print(cidx, total)

# print(blocks)
# print(blocks)
# print(out)
# t = 0
# for i, v in enumerate(out.strip("-").split("-")):
#     t += int(i) * int(v)
# print(t)
# t = 0
# for i, v in enumerate("00992111777.44.333....5555.6666.....8888.."):
#     if v != ".":
#         t += int(i) * int(v)
# print(t)
