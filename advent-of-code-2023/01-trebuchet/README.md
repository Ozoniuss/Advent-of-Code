# Day 1

This was definitely the most annoying day 1 I've tried. First part took me about 15 minutes and for the second one it took me over 1 hour and a half to spot the stupid mistakes I've been making.

Let me just quickly go through the [original commit](https://github.com/Ozoniuss/Advent-of-Code/commit/228022cfc1e25d47a9e7711ae602befba8abae17) and spot where I could have done better:

-   First, it took me an unbelievably long amount of time just to write part 1. I'm definitely not used to writing code so fast, but I have a feeling it will get better in time.
-   Lately I've been doing an approach similar to this one when parsing the input, because it's more efficient than using strings.Split()

```go
for i := 0; i < len(line); i++ {
	if line[i] >= '0' && line[i] <= '9' {
		if first == -1 {
			first = int(line[i] - '0')
		}
		last = int(line[i] - '0')
	}
}
```

But in this case there's no special parsing, you just read the entire line so that made me slower. I should have definitely thought of `strings.Index()` and `strings.LastIndex()`. They would have solved the problem for me, and part 2 would have followed naturally. I've updated the solution to use those functions.

-   I realized how easy it is to make dumb mistakes when optimizing for speed so hard. Currently the pace is more than I can handle. For example this line is so bad:

```go
total, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
```

It crossed my mind to multiply by 10^x, but I figured I can't know x. Until a good while later I realized that they're digits...

I made a lot of small errors while solving this. They were mostly around not using `strings.Index` correctly. In a string like `twothree`, both the first and second substring are relevant. For the first digit you need the left part and for the second digit you need the right part. I realized very late that once I switched to using the standard library substring finder, I was only computing the left ones.

This took much longer than it should have, but practice makes perfect. Let's keep improving that pace!
