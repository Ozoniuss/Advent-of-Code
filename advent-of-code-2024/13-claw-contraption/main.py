from collections import Counter
import re
from z3 import Solver, Int, sat, unsat
import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

equations = []
for i in range(0, len(lines), 4):
    equations.append((lines[i], lines[i+1], lines[i+2]))

params = []
for eq in equations:
    p = []
    print(eq)
    for expr in eq:
        print(expr)
        vals = re.findall(r'\d+', expr)
        p.append(list(map(int, vals)))
    params.append(p)

t = 0

# Assumes every equation has a unique solution or is unsolvable over
# the set of integers.
for param in params:
    s = Solver()

    a = Int('a')
    b = Int('b')


    tokens = 0
    # s.add(a * param[0][0]+ b * param[1][0] == param[2][0])
    # s.add(a * param[0][1]+ b * param[1][1] == param[2][1])
    s.add(a * param[0][0]+ b * param[1][0] == 10000000000000 + param[2][0])
    s.add(a * param[0][1]+ b * param[1][1] == 10000000000000 + param[2][1])
    if s.check() == sat:
        tokens += 3 * int(s.model()[a].as_string()) + int(s.model()[b].as_string())
    t += tokens
    print(param, tokens)

print(t)

# s = Solver()

# s.add((a * 1 + b * 1) == 10)
# s.add((a * 2 + b * 2) == 20)

# print (s.check())
# m = s.model()
# print(m)
# m = s.model()
# s.next()
# print(m.evaluate())
# # print(s.solve)