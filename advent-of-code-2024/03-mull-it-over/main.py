from collections import Counter

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

print("2".isnumeric())

# definitely not pretty but I pretty much only worked on this while waiting
# for pipelines to finish

t = 0
for line in lines:
    parts = line.split("mul")
    for part in parts:
        s = part.find("(")
        e = part.find(")")
        if s !=0 or e == -1:
            continue
        nums = part[s+1:e].split(",")
        if len(nums) < 2:
            continue

        firststr, secondstr = nums[0], nums[1]

        if firststr.isnumeric() and secondstr.isnumeric():
            t += int(firststr) * int(secondstr)

print(t)

t = 0
consider = True
for line in lines:
    parts = line.split("mul")
    for part in parts:

        dopos = part.rfind("do()")
        dontpos = part.rfind("don't()")
        # print(dopos, dontpos)

        changeatend = False
        toChange = None
        if dopos != -1 or dontpos != -1:
            # print("would change", dopos, dontpos)
            changeatend = True
            if dopos < dontpos:
                toChange = False
            else:
                toChange = True

        if consider:
            # print('skipping', part)
        
            skip = False
            s = part.find("(")
            e = part.find(")")
            if s !=0 or e == -1:
                skip = True
            nums = part[s+1:e].split(",")
            if len(nums) < 2:
                skip = True

            if not skip:
                firststr, secondstr = nums[0], nums[1]
                if firststr.isnumeric() and secondstr.isnumeric():
                    # print(part, int(firststr), int(secondstr))
                    t += int(firststr) * int(secondstr)
        
        if changeatend:
            print("changed to", toChange, dopos, dontpos)
            consider = toChange

print(t)