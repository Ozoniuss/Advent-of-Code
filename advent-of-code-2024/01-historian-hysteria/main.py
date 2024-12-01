from collections import Counter

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

left = []
right = []
for line in lines:
    p = line.split("   ")
    left.append(int(p[0]))
    right.append(int(p[1]))

t = 0
for k, v in zip(sorted(left), sorted(right)):
    t += abs(k-v)

print(t)

t = 0
c = dict(Counter(right))
for l in left:
    if l in c:
        t += l * c[l]
print(t)