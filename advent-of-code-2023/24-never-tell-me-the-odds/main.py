from z3 import Int, set_param, Solver
import math
import re

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

reg = re.compile("-?\d+")
vals = []
for line in lines:
    vals.append([int(num) for num in reg.findall(line)])


"""
This problem could have been nice, but instead it ended up being retarded. Maybe
there is a better solution to this but if the official solution requires plugging
into Z3 then part b is complete garbage.

However, note that my parallel line check criteria could have overflowed, so
in fact there might actually be parallel lines and what I said previously is
wrong. I'll check with python too just to be sure.

Some thoughts:

Assume there is a solution to this motherfucker. Then, if we project all
lines on the plane perpendicular to the line that intersects them all, all
intersections will project to a single point. As far as I'm aware, for that many
lines, that plane should be unique (especially since the statement mentions
there's only one line that intersects them all).

We're asked for the coordinates of point "0", but there are infinitely many 
possible initial points. This means that the vector of this line's coordinates
add up to 0. If we look at the example, it holds true (-3,1,2). So to begin
with I assumed the coordinates add up to 0.

Now, a nice trick here from a design perspective would have been to include at
least two parallel lines here. In my Go code I did not find any parallel lines,
so I ended up brute forcing this motherfucker. But, if there are two parallel
lines, they will include the line we're looking for, thus then the plane formed
by them lines is perpendicular to the plane that is perpendicular to the
required line. So we would be able to determine the "normal vector" of that plane
which is the direction of the line we're looking for. Its coordinates should
add up to 0. However it doesn't matter as much, just for validation

Now we can project some lines onto the plane and find their intersection. And
we have a point, so we just need to sum up its coordinates.
"""

a0 = Int("a0")
b0 = Int("b0")
c0 = Int("c0")

a = Int("a")
b = Int("b")
c = Int("c")

solver = Solver()

print(3 // -2)

# Disgusting solution, disgusting part b. If there would have been at least two
# parallel planes (see comment in Go file)
for idx, val in enumerate(vals):
    t1 = Int(f"t{idx}")

    x0, y0, z0, x, y, z = val

    solver.add(x0 + t1 * x - a0 - t1 * a == 0)
    solver.add(y0 + t1 * y - b0 - t1 * b == 0)
    solver.add(z0 + t1 * z - c0 - t1 * c == 0)

# set_param(verbose=2)
print(solver)
print(solver.check())
print(solver.model())
a0val = solver.model().eval(a0).as_long()
b0val = solver.model().eval(b0).as_long()
c0val = solver.model().eval(c0).as_long()

print(a0val + b0val + c0val)
