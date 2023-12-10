import sys
from functools import reduce
from dataclasses import dataclass
from collections import Counter
from functools import cmp_to_key

solvepart = None
if len(sys.argv) >= 2 and not sys.argv[1] == "--no-ask":
    solvepart = input("which part to solve> ")

f = open("example.txt", "r")
data = f.read().strip()
lines = data.split("\n")


@dataclass
class HandWithValue:
    hand: str
    value: int


hands_with_values = list(
    map(
        lambda line: HandWithValue(
            hand=line.split(" ")[0], value=int(line.split(" ")[1])
        ),
        lines,
    )
)
cardValues = {v: i + 2 for i, v in enumerate("23456789TJQKA")}


# I originally returned an array instead of a single value because I thought
# that ties were broken with a subsequent smaller combination but a score is
# enough.
def getHandScore(hand: str):
    int

    cards = dict(Counter(hand))
    if len(cards) == 1:
        return 7
    if len(cards) == 2:
        first = list(cards.values())[0]
        if first == 4 or first == 1:
            return 6
        else:
            return 5

    if len(cards) == 3:
        values = sorted(cards.values(), reverse=True)
        if values[0] == 3:
            return 4
        if values[0] == 2:
            return 3

    if len(cards) == 4:
        return 2
    return 1


def compareHands(hand1, hand2):
    if getHandScore(hand1) > getHandScore(hand2):
        return 1

    if getHandScore(hand1) < getHandScore(hand2):
        return -1

    for c1, c2 in zip(hand1, hand2):
        if cardValues[c1] > cardValues[c2]:
            return 1
        if cardValues[c1] < cardValues[c2]:
            return -1
    return 0


def compareHandsWithScores(hs1: HandWithValue, hs2: HandWithValue):
    return compareHands(hs1.hand, hs2.hand)


def part1():
    hands_with_values_sorted = sorted(
        hands_with_values, key=cmp_to_key(compareHandsWithScores)
    )
    values = [(idx + 1) * hs.value for idx, hs in enumerate(hands_with_values_sorted)]
    print(sum(values))


def part2():
    ...


if solvepart == None:
    part1()
    part2()
    exit(0)

ans = part1() if solvepart == "1" else part2()
print(ans)
