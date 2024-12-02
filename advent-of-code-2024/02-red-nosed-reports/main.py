from collections import Counter

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

def is_safe_2(line:str):
    nums = list(map(int, line.split(" ")))
    if not (nums == sorted(nums) or nums == sorted(nums, reverse=True)):
        return False
    for i in range(1, len(nums)):
        if not 1<=abs(nums[i]-nums[i-1])<=3:
            return False
    return True

def is_safe(nums: list[int]):
    
    diff = int(nums[1]) - int(nums[0])
    if diff == 0 or diff > 3 or diff < -3:
        return False

    initial_diff = diff

    for i in range(2, len(nums)):
        diff = int(nums[i]) - int(nums[i-1])
        print(diff, end=" ")
        if initial_diff > 0 and (diff < 1 or diff > 3):
            return False
        if initial_diff < 0 and (diff > -1 or diff < -3):
            return False
    print()
    
    return True    


def is_safe_tolerant(nums: list[int]):
    to_check = [nums]
    for i in range(len(nums)):
        to_check.append(nums[:i] + nums[i+1:])
    return any(list(map(is_safe, to_check)))

safec = 0
safet = 0
for l in lines:
    nums = list(map(int, l.split(" ")))
    if is_safe(nums):
        safec += 1
    if is_safe_tolerant(nums):
        safet += 1

print(safec, safet)