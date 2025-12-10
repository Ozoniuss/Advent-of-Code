from z3 import Int, Solver, Sum, sat, Optimize
from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

totals = 0
for idxl,l in enumerate(lines):
    parts = l.split(" ")

    sequencestr = parts[len(parts)-1][1:len(parts[len(parts)-1])-1]
    sequencestr=sequencestr.split(",")
    sequence = list(map(int,sequencestr))
    sequence = tuple(sequence)
    print("seq",sequence)

    buttons = []
    for i in range(1, len(parts)-1):
        p = parts[i][1:len(parts[i])-1]
        pp = p.split(",")
        buttons.append(list(map(int, pp)))
    print("butt", buttons)

    # each button is pressed bi times
    solver = Optimize()
    coef = [[] for i in range(len(sequence))]
    variables = []
    varnames = {}
    for i in range(len(buttons)):
        bts = buttons[i]
        for key in bts:
            varname = "b" + str(idxl) + str(i)
            variable = Int(varname)
            coef[key].append(variable)

            if varname not in varnames:
                varnames[varname] = None
                variables.append(variable)
                solver.add(variable >= 0)
        # variables.append(Int(" x"+str(idxl)+str(i)))
    print("coef", coef)
    print("vars", variables)

    toMinimize = solver.minimize(Sum(variables))

    for idx, c in enumerate(coef):
        solver.add(Sum(c) == sequence[idx])

    mins = 0
    if solver.check() == sat:
        m = solver.model()
        print(m)
        # for c in m:
            # print("c", m[c])

        tot = m.evaluate(Sum(variables))
        print("eval", tot)
        totals += int(tot.as_long())
    else:
        print("No solution")

    print("mins", mins)
    # print(solver.model())


   

print(totals)



# part 1
# totalIt = []
# for l in lines:
#     parts = l.split(" ")

#     sequencestr = parts[0][1:len(parts[0])-1]
#     sequence = []
#     for c in sequencestr:
#         if c == '.':
#             sequence.append(0)
#         else:
#             sequence.append(1)
#     print("seq",sequence)
#     sequence = tuple(sequence)

#     buttons = []
#     for i in range(1, len(parts)-1):
#         p = parts[i][1:len(parts[i])-1]
#         pp = p.split(",")
#         buttons.append(list(map(int, pp)))
#     print("butt", buttons)

#     currentSeq = tuple([0] * len(sequence))
#     q = [currentSeq]
#     length = 0
#     found = True
#     while found:
#         nextq = []
#         for el in q:
#             # print("el", el)
#             for button in buttons:
#                 nextel = list(el)
#                 for key in button:
#                     nextel[key] = (nextel[key] + 1) % 2
#                 nextel = tuple(nextel)
#                 if nextel == sequence:
#                     found = False
#                     break
#                 if nextel not in nextq:
#                     nextq.append(nextel)
#         length += 1
#         q = nextq
#     totalIt.append(length)

# print(sum(totalIt))
