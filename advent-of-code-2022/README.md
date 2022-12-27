# Advent of Code 2022

Welcome! 🖖

This year, I decided to attempt the [2022 Advent of Code](https://adventofcode.com/2022) contest. I've heard about it from a friend, who basically told me that one problem gets published every day for 25 days and that they're not too difficult. I've never really been into competitive programming, but I was always attracted to problem solving, which I didn't really have that much time for. The idea of the contest was really intriguing to be, and after reading the first two problems I decided to challenge myself to solve all problems during the 25 days, every day after work and university.

This has turned out to be quite a ride. At first, the problems were fairly easy, and I just had to get used to writing the algorithms. I went with Go as the programming language, since I'm a Go programmer. Once I got in shape by solving the first two or three problems, it took me 20-30 minutes to solve the easier ones, but the difficulty kept increasing, and I even ended up spending hours (sometimes even in different days) on some problems. I must say, sometimes after the 16th day, not exactly sure which one, there were times when I felt it was a bit overwhelming and that I was spending too much time on the problems, and was wondering if I was going to be able to complete all of them. 

In the end, I believe this was very beneficial to me, and even though it cost me some hangouts with friends and Friday nights (I think somewhere around 30 hours of work to solve all these problems), I'm glad I did it. I feel that I have much more control over the fundamental algorithms, and my code writing speed has increased significantly. I was also able to improve my Go skills via this challenge, so every hour spend past 12 a.m. to grind these problems was worth it. It brought back the joy of getting a problem right back from the old college days when I was doing math contests, and I have to admit that the last day when I finally got the only problem I had remaining right, problem 19, I was so delighted 🥳! I would gladly recommend this contest to anybody, really (although some prior basic algorithms and data structures knowledge would be an enormous benefit).

On a related note, I never really played it just for speed. Increasing my coding speed was one of my main goals, but at all times I tried to keep the code clean, readable and have a good separation of concerns. I might have slacked out sometimes when I spent a lot of time on the problem because I was happy to just get it done, but I'll likely refactor the code I don't fancy. There will be a link to the original solutions anyway for each problem.
 
Below you can find a _very simplified_ summary of all the contest's problems. The folders contain my own solutions to these problems, and each readme has some notes I found interesting. I did every problem myself, without looking for any hints or inspirations (aside looking up some basic techniques like backtracking to refresh the concepts), so the solutions may not be optimal, but I'm happy all of them ran in a good amount of time.

### Enjoy!

---

1. [🍩🍫🍪](./01-calorie-counting/): intro

This problem was pretty easy, the main idea bringing down to computing the first, and first 3 elements of an array. Maybe doing the second part at the same time of reading the input could have a touch of trickiness to it, but there was nothing really outstanding about this one.

I would give this one a difficulty score of one Christmas Tree: 🎄

2. [🧱📄✂️](./02-rock-paper-scissors/): rock, paper... scissors!

Like the title says, here you had to simulate a rock-paper-scissors game. The first part was giving you the exact moves for several games and you had to compute your score, while in the second part the same input actually represented the outcome of the games; you had to find out what you needed to play for each game and then compute the score for all games based on that. There wasn't really anything difficult about this one either, perhaps a nice thought exercise would be figuring out the simplest way of modelling the rules of the game, like what beats what.

The difficulty score is still one Christmas Tree: 🎄

3. [💼📛📁](./03-rucksack-reogranization/): the common item

The first part of this problem is pretty much taking a string with even length, splitting it exactly in half, and finding the guaranteed unique common letter. There are straightforward ways to do this, by just checking which letter from one half is in the letters of the other half, but this becomes more interesting once you try to "optimize" this (as if you need to for 50-character strings 😉)

Of course, me being a huge nerd I did optimize it, by considering each letter being a bit (a _bit_, yes, not a _byte_; the entire alphabet has 52 characters and that fits in a 7 byte number) and finding the intersection via bitwise operations, `and` to be specific. You can view my madness 😵 [here.](./03-rucksack-reogranization/main.go) 

The only difference in the second part is that you're intersecting three sets instead of two. It's more or less the same idea, you just check that the current item is in the other two sets as well. With my completely useless optimization, you just do 3 bitwise `and` operations instead of 2.

Perhaps this was not as straightforward as the other two, especially if you wanted to optimize a bit the set intersection, but you can still just get the problem done pretty fast. For this reason, I consider the overall difficulty of this one to still be one Christmas Tree: 🎄

4. [➕♾️➰](./04-camp-cleanup/): sketchy intervals

Time for some more math-y bussiness. With problem 4, you had to do some work with intervals. The first part was figuring out from two intervals if one was part of the other, whereas the second part was figuring out if the intervals intersect. Since the intervals consisted only of consecutive integers, you could just go through all the numbers, but you can also answer both questions just by looking at the interval's bounds. Perhaps figuring out the conditions with the bounds instantly is a bit trickier, but it doesn't take long to get them right.

For this reason, this was one of the shortest problems and it also gets one Christmas Tree: 🎄

5. [🚢📦🚧](./05-supply-stacks/): moving boxes

This is the first (if you don't count rock, paper, scissors) out of many simulation problems, and it simulates a crane moving boxes from a pile to another out of many piles. It's also the first one where the input wasn't straightforward to parse; in fact, parsing out the boxes was slightly challenging:

```
[D]                     [N] [F]    
[H] [F]             [L] [J] [H]    
[R] [H]             [F] [V] [G] [H]
[Z] [Q]         [Z] [W] [L] [J] [B]
[S] [W] [H]     [B] [H] [D] [C] [M]
[P] [R] [S] [G] [J] [J] [W] [Z] [V]
[W] [B] [V] [F] [G] [T] [T] [T] [P]
[Q] [V] [C] [H] [P] [Q] [Z] [D] [W]
 1   2   3   4   5   6   7   8   9 
```

This is probably the point where I should have started using regex to read the input file, they made my life much easier down the line once I added them. It would have definitely made reading the moves nicer too:

```
move 1 from 3 to 9
```

Moving on to the simulation, at every step you move a number of crates from one pile to the other. Part (b) is the exact same, except that it reverses the order of the crates at every move. It's also guaranteed that every move is valid, that is, there's always enough crates in the pile you're taking them from.

The simulation itself was not challenging, but storing state does introduce an additional layer of complexity, especially when you're trying to write clean code. This, combined with parsing the input makes this the first problem that receives from me two Christmas Trees in difficulty: 🎄🎄

6. [📱❎🈶](./06-tuning-trouble/): finding the right substring

Back to string work. The input was actually only a long string this time. The main idea behind this problem was finding the first substring (with consecutive characters) with a given length from the long string such that all the characters in the substring are different. In part (a), that lenght was 4, and in part (b) that length was 14, probably to add some more algorithmic thinking to those who at part (a) just checked that each letter is different from the other 3 (yes, myself included 🙈).

Either way, this problem was really simple, slightly similar underlying idea to problem 3 where you had to check for a common character. The most straightforward ideas would be converting the substring to a set, or checking char by char in two `for` loops. Doing that in a function while advancing with the substring one letter at a time through the main string also makes for a nice separation of concerns. Technically you could also optimize by keeping track of the characters in the previous check, but that's not really necessary.

```
dcbc(sbblhhgdgssmcm)qccdw  -->  dcbcs(bblhhgdgssmcmq)ccdw  --> ...
```

This problem was quite a bit shorter than the previous one, so it only gets one Christmas Tree in difficulty: 🎄

7. [📁💾🖥️](./07-no-space-left-on-device/): cleaning the filesystem

Somewhat of a simulation as well, this problem is inspired from the Linux filesystem and in my opinion is really well thought-out. Your input consists of two Linux commands: `cd` (change directory) and `ls` (list). The problem simulates navigating through the filesystem and displaying the files from various directories, including their size. In both parts, the requirement was pretty much computing the size of each directory, and identifying the one(s) that satisfies a specific constraint, in particular, having a size greater than some number at part (a), and being the smallest one with the size greater than some number at part (b).

Once you know the size of all directories, answering the questions is not hard at all, but the tricky part here is to actually find the size of each directory! A stack is your best friend to keep track of the current directory. If you know your working directory, you can pretty much interpret the meaning of every line, so this problem can be solved by reading the input line by line and holding the current directory and its size. In the input it's also pretty much guaranteed that the exploration pattern is always the same.

Even if the idea itself doesn't seem complicated, the execution is not trivial. One must take into account all parent directories for a directory name, since I didn't find the statement to specify whether directory names are unique, and must also keep in mind that to the size of a directory one must add the size of the child directories. With all these in mind, this "simulation" was definitely more challenging than the crane one for problem 5, giving the problem three Christmas Trees in difficulty: 🎄🎄🎄

8. [🌲🏡🚁](./08-treetop-tree-house/): the tallest tree

This is the first problem involving a matrix: a matrix of integers. This is also the first problem where I started benefiting from the fact that in Go, arrays are passed by value by default. Whenever I had to store positions, I went with the datatype `[2]int`, and later down the line, I started doing 

```go
type Location [2]int
```

as syntactic sugar.

To break down the requirement, at part (a) you had to find all numbers in the matrix that's either greater than all numbers in front or behind it, either on its row or column. Part (b) was similar, for all numbers in the matrix you had to compute a property based on how many numbers on its row or column are greater than itself, and find the position with the highest value for that property.

There likely are better optimizations, but the most intuitive approach of going in all four directions using loops does the trick for this problem. And since that and reading a matrix input is straightforward, I will give this problem one Christmas Tree in difficulty: 🎄

9. [🐍🌉🧵](./09-rope-bridge/): snake

On with the simulations: this problem was pretty much simulating the snake game (or at least a very similar game). For part (a), the snake had length 2 and length 10 at part (b). The statement describes how the snake "moves", and the input consisted of the snake directions. In both parts, you had to say which squares the snake visited.

Using the same trick as before for positions and modelling the snake as an array of fixed length (in golang) works well here, and defining a `move(snake, direction)` function does a nice separation of concerns. Once the modelling is done, part (a) is done pretty quickly, but applying the same idea for part (b) doesn't work, because the motion can place the snake in certain positions that don't come to mind initially, where the movement is different. These were also not showcased in the example, so I did spend quite some time wondering why the approach at part (a) kept failing. I'll describe that case in more detail on the problem page, but for that reason, this problem receives three Christmas Trees in difficulty: 🎄🎄🎄

10. [🕹️📺🔍](./10-cathode-ray-tube/): hardware instructions with a CRT screen

The simulations keep going, now simulating cycles of a circuit and drawing pixels. At part (a), the instructions were simulating adding and subtracting values to a register at every cycle, and you had to keep track of the register's values at some specific cycle numbers. Part (a) went quickly, but part (b) was definitely a ramp up in difficulty: the cycle number was simulating a pixel position on a CRT screen, and based on the register's value (which indicated three consecutive pixel positions) you had to determine whether you were allowed to draw a pixel or not. The fact that after 40 pixels you started a new row also added to the difficulty.

This was also the first and only problem that required a human interpretation of the result: at the end you had to draw the "screen", and write the letters displyed on the "screen" as the answer. It definitely was one of the longer problems, but the checks themselves were not difficult, maybe a bit annoying because you had to keep track of multiple things at once. Talking about difficulty, I would consider it on the same page with the previous one, receiving three Christmas Trees: 🎄🎄🎄

11. [🐒⚾🏓](./11-monkey-in-the-middle/): _monke_ games

This is where the monkeys (or _monke_-s) got in the picutre. In this new simulation problem, at each round, the monkeys have a list of numbers, and each monkey does some operations to every number. Based on the result of the operation, the monkey decides which monkey to send the new number to. The requirement was to compute how many numbers were thrown away by each monkey after a number of rounds.

The difference between part (a) and part (b) were the operations with the numbers. The results of the operations at part (a) were reasonably small numbers, but at part (b) you would very easily start to get huge numbers that didn't even fit to 64 bits fairly quickly, so you had to come up with a trick to keep the numbers under control. Figuring out why the approach from part (a) is failing and the trick to avoid that is not straightforward and might take some time, thus awarding the _monke_-s with three Christmas Trees in difficulty: 🎄🎄🎄

12. [🏁🌳⛵](./12-hill-climbing-algorithm/): shortest path

We knew this had to come eventually: the first backtracking problem! This one is quite a classic though: you start from one place in a matrix and you have to find the shortest path to a different place, with the condition that the neighbours you can go to from some point are based on the letter they have. Part (b) is not that much different, except that there are multiple starting points and you must additionaly find the starting point which gives the shortest path.

The idea to solve this problem is pretty well-known, and that is, [bread-first search](https://www.geeksforgeeks.org/shortest-path-unweighted-gr) (commonly referred to as `BFS`). The idea for part (b) is almost identical, known as [multi-source bfs](https://www.geeksforgeeks.org/multi-source-shortest-path-in-unweighted-graph/). This problem can be done quickly by defining a `getNeighbours(Location) -> Location` and applying these well-known techniques, but because backtracking is a harder concept to graps in general, the problem receives two Christmas Trees in difficulty: 🎄🎄

13. [🔃↔️🔄](./13-distress-signal/): infinite recursion

This problem introduced perhaps one of the most difficult concepts to define formally, and that is, recursive objects. In particular, lists whose items can be either integers, or lists of integers, even emtpy lists. The problem's input is a bunch of pairs of these types of lists, and part (a) requires you to "compare" the lists, based on a recursive method defined by the statement. Part (b) is much easier assuming you got part (a) right; it only requires writing all input lists, together with two additional ones, in the "correct" order based on the comparison method.

There are several tricky parts to this method, that is, defining the recursive object that holds the lists and numbers, and converting between that object and its string representation. Once that's done, implementing the comparison function on the recursive object is less of a challenge since it's clearly defined in the statement, but the string convertion is quite brainy and for this reason, this problem is the first one I decided to give four Christmas Trees in difficulty: 🎄🎄🎄🎄

---

I believe this was the point in the contest where I knew I would be in for quite a ride. Last two problems introduced two of the harder concepts in computer science, backtracking and recursion, and problems started to take increasingly longer to solve. At times I was wondering whether it was the right moment for me to do this challenge, since I often ended up working on these problems after midnight, but this last problem was simply beautiful and I decided to keep going nonetheless.

---

14. [⏳🕰️⛰️](./14-regolith-reservoir/): falling sand

The simulations keep coming, this time sand falling down a map which also contains rocks in various places. The sand falls from some location and when it encounters a rock (or other sand) it can either stop or fall to the left or right. The puzzle input represents the rock patterns on the map, and one unit for sand falls at each step. Part (a) requires to compute the quantity of sand that falls until all new sand units will no longer come to rest and fall down the map (which is guaranteed to happen), where in part (b) a straight rock "border" is placed at the bottom such that sand never falls down, and you have to find when the sand source gets blocked by the pile of sand.

This problem required a fair amount of work: the input is not straightforward to parse, adding the border at part (b) is a bit cumbersome, simulating the sand motion did require a bit of recursivity and due to the nature of the problem, erros are more difficult to debug (you can draw the map, but you have to compute the region of the map you want to draw first, and then actually draw it...). However, the core idea wasn't really that difficult, especially compared to the previous problem and therefore collecting 3 Christmas Trees in difficulty: 🎄🎄🎄

15. [💎🛑🚧](./15-beacon-exclusion-zone/): map coverage

This next problem is placed in the 2-dimensional plane. It provides the integer coordinates of a few sensors as input, each one being able to scan all points up to a given distance (using Manhattan distance). Part (a) gives a line and asks how many integers points located on that line are detected by these sensors. At part (b), you were given some delimited area where you knew that exactly one point is not scanned by the sensors, and asked to find that point.

With an intuitive approach, one could write a function `isVisible(location, sensor) -> bool` which will get part (a) done, but would be slow for part (b). The catch here is that coordinates are massive: you can compute the manhattan distance instantly but there are multiple sensors and the area at part (b) is 4 million by 4 million, so just checking each individual point is too slow, at least in Golang. Because it did require an optimization, the sensor's coordinates were of the order of millions and computing the bounds of the line at part (a) is slightly trickier (after all you can't check the entire line), this problem covers 3 Christmas Trees in difficulty: 🎄🎄🎄

Btw, if you're parsing the input with regex, watch out for negative numbers. That did cost me one hour 😉

16. [🗻🎡🌋](./16-proboscidea-volcanium/): best-scoring road

The contests introduced again a backtracking problem in day 16. This time there were a bunch of valves and tunnels between them (so basically a graph with edges of length 1) and opening each valve would release a certain amount of pressure each minute (for some valves, that could be 0). At every minute, you could either open a valve or move one edge to a different valve, the requirement being to maximize the released pressure after 30 minutes. Part (b) was pretty much the same, except that you had 26 minutes and a "companion" to open valves with you, that could do the exact same operations.

Unlike week 12, this was a less straight-forward approach of backtracking with DFS (or BFS). One had to construct the graph from an input that wasn't trivial to parse, there were some challenges to model the steps nicely because it's actually harder to store the state minute by minute, and coming with an approach for part (b) does require some thought. The implementation definitely takes some time, and combined with the difficulty made the problem worthy of 4 Christmas Trees: 🎄🎄🎄🎄

---

On a more personal note, this is the first problem that discouraged me because I wasn't able to finish it the day I started it, and neither the next problem which was even harder. I actually completed it on day 21 I think, when I solved 3 problems. I kept doing the really stupid mistake of adding the starting point twice to the input graph at part (b), which kept messing up my final answer. When I realised what the issue was, my reaction was best described by 😒

This is also the first problem I optimized with parallelism, which decreased the time to solve part (b) substantially. I will definitely be looking to add more parallel programming when remastering these problems.

---

17.  [🧱🔷🟩](./17-pyroclastic-flow/): Tetris (or at least something close)

Around the corner was sitting likely the problem that I liked the most. In this next simulation, you had to simulate a game of Tetris, where you were given 5 shapes and a sequence of left or right keyboard inputs. Both the sequence of shapes and keyboard inputs was repeated once it ended, and after each keyboard input the piece dropped exactly one unit. Parts (a) and (b) were the exact same task: to compute the height of the tower after 2022, and 1 trillion rocks fell, respectively.

Yeah, 1 trillion -- you heard that right. Simulating the entire game was possible at part (a) and the approach I went with, but that just wasn't going to cut it for part g(b). The only hope there was finding some patterns. And finding a pattern that is proven mathematically to work is very difficult: what is the condition for the pattern to repeat? There were also 5 different pieces, and 10091 keyboard inputs (in my puzzle input at least). Not only that, but even once you found the pattern, the math behind required to compute the height using the pattern is simply insane.

After solving part (a) fairly quickly, the difficulty increase brought by part (b) was baffling to me. Adding everything together: modelling the input and states, finding the pattern (keep in mind that one can prove in some games there is _no cyclic pattern_) and computing the height based on the pattern, in my opinion completely justifies the maximum number of Christmas Trees I'm giving this problem in difficulty: 🎄🎄🎄🎄🎄

18. [🧊🔥🌪️](./18-boiling-boulders/) structure of mini-cubes

Breaking the monotony of the 2-dimensional space is, in my opinion, one of the most elegant problem of the contest. This one pretty much consisted of placing unit-cubes in the 3-dimensional space. The input was giving the positions of a bunch of `1x1x1` cubes, and at part (a) you had to compute how many faces of these cubes are not touching another cube. Part (b) did make matters more interesting, and the easiest way to visualize the requirement for me is to think that the generated cube structure is plunged into water, and you had to count the number of faces that get wet.

While I managed to get part (a) done really quickly, part (b) brought additional layers of complexity, and did require using your spatial orientation a lot when designing an algorithm to solve it. It the end, still an exploration problem solved by BFS, but understanding what you had to do, visualizing the "hidden" spots not reached by water and coming up with a proven working algorithm is definitely challenging. Generally it is much harder to work with points in space than in the plane, but I'd say the problem was overall a bit easier than the previous one, making 4 Christmas Trees a fair difficulty score: 🎄🎄🎄🎄

19. [🤖⛏️⛑️](./19-not-enough-minerals/) and yet another backtracking problem

I'll just start by saying that this problem did make me want to rip my hair off my head. It is the last problem I submitted, which got me the 50th star. It is a bit similar to finding the best investment: you were given 4 different robots, each being able to gather one type of resource (ore, clay, obsidian, geode) every minute. You were also given multiple blueprints with the cost of each robot, and you had to find out the best strategy to buy robots when you had the resources in order to maximize the number of geodes after 24 minutes, for each blueprint. Part (b) is essentially the same as part (a), except you only had to compute the maximum number of geodes for the first three blueprints, but after 32 minutes.

The technique that solves this one is backtracking with DFS (or BFS but that takes a huge toll on memory) like we've seen before, but the catch is that is suffers terribly from the state-space explosion problem without some good pruning. And finding that good pruning is hard. Even after hashing each possible state and trying to estimate the maximum number of geodes from every state, it still took a good 10 seconds to run part (b), and it would have probably taken days without any pruning at all. I might also be a bit biased here due to solving it last, but because it was challenging to model the states (which for me included the number of robots, number of states and some action) and to find some good pruning then implement it, I consider this problem alongside the "Tetris" one from day 17 to be the only two ones worth 5 Christmas Trees in difficulty: 🎄🎄🎄🎄🎄

Just to strengthen why I think modelling this problem is so difficult, I only figured out on day 25 that my original approach was incorrect and allowed creating robots one round before they were actually created. After fixing this, the state explosion moved from minute 14 to minute 20, which was a significant improvement. I did however also find the statement itself to not reveal that detail clearly; I'll go into more details on the notes I make for the problem on GitHub.

20. [1️⃣♻️#️⃣](./20-grove-positioning-system/) cyclic list

The waters chill down a bit on day 20, after what I think were 3 of the most difficult problems of the contest, one after the other. The statement had much more simplicity, the input being just the elements of a list of integers. The list is considered to be cyclic, as if the elements were in a circle, and for part (a) you had to go through all the elements from the input, and for each one move it on the circle a number of positions equal to its value. So if the element was 3, you'd move it past the 3 next elements, if it was -2 you'd move it before the previous two elements. Part (b) was the same, except that the numbers were all multiplied by some large number like `811589153`, and you had to repeat the process by going through the original list of numbers 10 times. The task required to find the numbers on three positions.

Although this problem is nowhere near as difficult as the previous one, it still has its quirks. There's deciding on how to store the elements; if using a normal array, moving the elements has some edge cases, otherwise one could implement a linked list. Also, one should implement an optimization to avoid unnecessary round trips around the circle of numbers. And keep in mind that there are also duplicates, meaning that you have to know which number to move. Figuring all of these out increased the time I spent on this problem way past what I thought initially it would take me, for which reason I grant the problem three Christmas Trees in difficulty: 🎄🎄🎄

---

Fun fact: I originally got the answer wrong by using an array implementation, so I switched completely to using a linked list. I still got the same answer, though. It turned out that I changed my implementation completely just to find out that I had copied the answer from Microsoft's calculator with a comma: `27,726` instead of `27726` 🙄. Nevertheless, I much prefer the linked list approach, but I would have rather done to something else than spending one more hour on this at 2 a.m.

---

21. [〰️🔱🌿](./21-monkey-math/) doing math with recursion and binary search trees

Coding the solution to this problem was just nice. You were given a bunch of math expressions: on the left side, some variable. On the right side, either a math expression (+,-,/,*) with two different variables or some value. So basically you either knew the value of the variable, or had to compute it by going through the sub-expressions of the variables in the original variable's expression, recursively. You can see how this leads nicely to a binary tree of expressions. For part (a), you had to find the expression of the `root` variable, which was the top of the tree.

While part (a) was a simple recursion, part (b) was much more challenging. There was a variable called `humn` that had a value initially, but this time you had to specify which value it needs to have such that the two parts of the `root` expression have equal values. This did require coming up with a clever way to parse the tree to get an expression for `humn`, which was a leaf in the tree. Seeing the approach is not straightforward, and recursivity and binary trees are hard concepts, thus my expressions determined that four Christmans Trees is the value for the difficulty variable: 🎄🎄🎄🎄

22. [🦧🗺️🎪](./22-monkey-map/): expanded cubes

Incoming simulation alert: this time, an oddly shaped map and a set of moves and direction changes as input. Basically, you knew how many places to go forward before changing the direction to either left or right. For part (a), when you fell off the map, you spawned in the opposite location on the same line or column. There were also rocks that could obstruct your movement, even at the other end, when you tried to wrap across the map. Part (a) required to compute the final location after the moves were completed.

The contest did make us familiar with simulations, and thus part (a) was nothing out of the ordinary. At part (b) though, it turned out that the map is in fact the expanded version of a cube, and on the two-dimensional expansion you had to simulate moving on the cube's faces. This significantly ramped up the difficulty, because mapping the edges to one another was a pain, and the behaviour now depended on direction too around the corners. There was also a very subtle trap at part (b) you could fall into if you were interpreting directions in a particular way at part (a) that I'll be discussing on the problem's page, and for all these reasons this was amongst the harder simulations, paving the path to 4 Christmas Trees in difficulty for the problem: 🎄🎄🎄🎄

23 [🗺️🏔️🌱](./23-unstable-diffusion/) another map simulation??

Severe simulation warning ⚠️. However, this time an easier one. Your input is a bunch of elves' locations on an infinite map. At each step, you know a list of directions where each elf can move, and the elves can move in a direction if they aren't "too close" to neighbouring elves, or isolated from them. At part (a), you had to enclose the elves in a square with smallest area after simulating 10 moves, and had to count the empty squares. And at part (b), you simply had to run the simulation until all elves were isolated from their neighbours, answering with the number of rounds it took to reach there.

Overall, the movements and move conditions were pretty well explained and easy to model, but there were trickier parts, in particular, dealing with an infinite grid and addressing move conflicts. This type of problem is also not easy to debug, especially if the number of elves is large, therefore receiving a total number of three Christmas Trees in difficulty: 🎄🎄🎄

24 [➡️🚸⛈️](./24-blizzard-basin/) path finding on moving map

Last backtracking problem. Well, at least it was a challenging one. The input is a bounded map in the shape of a rectangle, a starting point and a finish point. On the map, there are also a big number of winds going in one of four directions, passing through each other and wrapping around the map when they reach the bounds. You just have to find the shortest path from start to finish, with the constraint that you must never be on the same tile with a wind. Part (b) is basically the same task, repeated three times: you need to find the shortest path from start to finish, then finish to start, and then back again start to finish (there will be different paths because winds will be in different positions).

Of course, the main technique is still backtracking with BFS, but there's the additional difficulty of keeping track of the wind positions at each step, actually moving the winds, and remembering to move if a wind takes your place (it is guaranteed that you have a valid place to move). Moreover, you will need some pruning techniques, otherwise your algorithm will run in hours (at least in Go). Thankfully, exploring (step, location) pairs reached via different paths only once is a good enough pruning which allows BFS to complete in a few minutes, but the difficulty is still there and there was quite a lot of work to simulate those winds, making this problem one step above the previous simulation in difficulty at four Christmas Trees: 🎄🎄🎄🎄

25 [❓➗5️⃣](./25-full-of-hot-air/) a different form of base 5 numbers

After a bumpy ride, the adventure finally comes to and end with, luckily for me, a math puzzle. The problem comes up with a different system to write numbers, the possible digits being `2`,`1`,`0`,`-`,`=` with values `2`,`1`,`0`,`-1`,`-2` respectively. To compute the value of a number written with these digits, simply multiply each number with 5 to the power of its position and add those together, for example, `2=1` is `2*5^2 - 2*5^1 + 1*5^0 = 50 - 10 + 1 = 41`. The problem had only one part: you were given a list of numbers using this representation, had to add them up and write the sum using this representation.

Luckily, there were no negative numbers, and all you really had to do was writing number modulo 5 in a slightly different way. Converting to base 10 is done in the same way, and once you replace the possible remainder values with the set `{-2,1,0,1,2}`, all there is left to do is re-define the quotient and remainder modulo 5 as `q = (x+2)/5` and `q = (x+2)%5-2`, the standard technique of converting to base 5 can be used for this new counting system. I likely found this really simple because of being a mathematician in the past, but I still think the conversions could be a bit tricky to get right, and one also needs to associate with the symbols of the new counting system, making the end of this contest marked by a nice two Christmas Tree difficult problem: 🎄🎄

---

Note: thinking about how to do the conversions for negative numbers gets interesting, because the "zero" values don't align: a negative number can start with 1 in the new counting system. I will likely cover that too on the problem's page.

---

That was the end of the 2022 Advent of Code contest. I had a tremendous amount of fun, and for those of you who participated, I hope you also enjoyed the journey! Overall, the puzzles and story were really well-thought, and I believe it took a lot of effort to generate such relevant puzzle inputs.