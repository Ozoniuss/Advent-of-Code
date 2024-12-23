from collections import Counter

from shlex import join
import sys
import os
from time import sleep

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

registers = {"A": 0, "B": 0, "C": 0}
isp = 0
output = ""


def combo(n):
    global registers
    if n < 4:
        return n
    if n == 4:
        return registers["A"]
    if n == 5:
        return registers["B"]
    if n == 6:
        return registers["C"]


def adv(n):
    global registers
    registers["A"] = registers["A"] // int(2 ** combo(n))


def bxl(n):
    global registers
    registers["B"] = registers["B"] ^ n


def bst(n):
    global registers
    registers["B"] = combo(n) % 8


def jnz(n):
    global registers
    global isp
    if registers["A"] == 0:
        return None
    isp = n
    return "ISP CHANGED"


def bxc(n):
    global registers
    registers["B"] = registers["B"] ^ registers["C"]


def out(n):
    global output
    output += str(combo(n) % 8) + ","


def bdv(n):
    global registers
    registers["B"] = registers["A"] // int(2 ** combo(n))


def cdv(n):
    global registers
    registers["C"] = registers["A"] // int(2 ** combo(n))


opc = [adv, bxl, bst, jnz, bxc, out, bdv, cdv]

program = None
for line in lines:
    parts = line.split(" ")
    if parts[0] == "Register":
        if parts[1] == "A:":
            registers["A"] = int(parts[2])
        if parts[1] == "B:":
            registers["B"] = int(parts[2])
        if parts[1] == "C:":
            registers["C"] = int(parts[2])
    if parts[0] == "Program:":
        program = list(map(int, parts[1].split(",")))

print(program, registers)
while isp < len(program):
    instr, op = program[isp], program[isp + 1]
    if opc[instr](op) is None:
        isp += 2
output = output.rstrip(",")
print(output)


programstr = ",".join(list(map(str, program)))
i = 0

# if input has binary form
# xxx xxx xxx xxx xxx xxx xxx xxx xxx
# then basically 000 represents 0 which is ignored, and the rest represent a
# specific value in the output. we need to find which value produces the
# program
# program has length 16, so there will be 16 such groupings, where the first
# one is nonzero


def base_8(n):
    if n == 0:
        return "0"
    out = []
    while n > 0:
        r = n % 8
        n = n // 8
        out.append(str(r))
    out.reverse()
    return "".join(out)


def from_base_8(n):
    digits = list(str(n))
    v = 0
    m = 1
    for d in reversed(digits):
        v += m * int(d)
        m *= 8
    return v


# print(base_8(7), from_base_8(7))
# print(base_8(8), from_base_8(10))
# print(base_8(16), from_base_8(20))
# print(base_8(50), from_base_8(62))
# os._exit(0)

# s = "abcde"
# print(s[-3:])
# os._exit(0)

"""
fn(x) -> random function  n returning n digits for x
f1(a) = f1(a)
f2(ab) = f1(ab) + f1(a)
f3(abc) = f1(abc) + f2(ab) = f1(abc) + f1(ab) + f1(a)
f4(abcd) = f1(abcd) + f3(abc)
...
we need f(16)(p) = p
so we make a mapping where for fn(ab...) we only store the values that are equal
to the last n digits of the program
"""


def debug():
    global isp
    global output
    i = 0
    while True:

        registers["A"] = i
        registers["B"] = 0
        registers["C"] = 0
        isp = 0
        output = ""
        # print(f'{registers["A"] : <8}', end=" ")
        print(base_8(registers["A"]).rjust(8, " "), end=" ")
        while isp < len(program):
            instr, op = program[isp], program[isp + 1]
            if opc[instr](op) is None:
                isp += 2
        output = output.rstrip(",")
        print(output)
        if output == programstr:
            break
        i += 1


# debug()
# os._exit(0)

programstr_nocomma = "".join(list(map(str, program)))
potential_values_for_prev_digit = [""]
ok = False

while True:

    if ok:
        break

    to_try = []
    for i in range(8):
        # if i == 0 and potential_values_for_prev_digit[0] != "":
        #     continue
        for potential in potential_values_for_prev_digit:
            to_try.append(potential + str(i))
    print("to try", to_try)

    next_potentials = []
    for candidate in to_try:
        number = from_base_8(candidate)
        l = len(candidate)  # the number of digits the output would have

        registers["A"] = number
        registers["B"] = 0
        registers["C"] = 0
        isp = 0
        output = ""
        # print(f'{registers["A"] : <8}', end=" ")
        # print(base_8(registers["A"]).rjust(8, " "), end=" ")
        while isp < len(program):
            instr, op = program[isp], program[isp + 1]
            if opc[instr](op) is None:
                isp += 2
        output = output.rstrip(",")
        output_nocomma = "".join(output.split(","))
        print(number, candidate, output_nocomma, programstr_nocomma)
        assert len(output_nocomma) == l

        if output_nocomma == programstr_nocomma[-l:]:
            next_potentials.append(candidate)

        print("next", next_potentials)
        if output == programstr:
            print("found it", number)
            ok = True
            break

    potential_values_for_prev_digit = next_potentials
