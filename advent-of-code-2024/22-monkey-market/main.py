from collections import Counter

import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

def mix(v, s):
    return v ^ s

def prune(s):
    return s % 16777216

def evolve(s):
    a =  prune(mix(64*s, s))
    b = prune(mix(a // 32, a))
    c = prune(mix(2048 * b, b))
    return c

buyers = []
for line in lines:
    buyers.append(int(line))

for i in range(2000):
    for j in range(len(buyers)):
        buyers[j] = evolve(buyers[j])

print(sum(buyers))

# s = 123
# for i in range(10):
    # print(s)
    # s = evolve(s)