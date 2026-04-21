# printreversealphabet

## Required

- Type: required
- File to submit: printreversealphabet/main.go
- Allowed function: github.com/01-edu/z01.PrintRune

## What This Exercise Is Teaching You

This exercise is a direct follow-up to printing the alphabet in normal order.

Now you are learning how to:

1. print characters in reverse order,
2. control loop direction by counting backward,
3. use rune ranges correctly from z to a,
4. follow strict constraints like no casting.

This is a very important beginner step because many programs need reverse loops.

## Task In Simple Words

Write a program that prints:

zyxwvutsrqponmlkjihgfedcba

on one line, then ends the line with newline.

## Output Rule You Must Respect

The output should be one line of letters, followed by end-of-line.

Expected terminal view:

```text
zyxwvutsrqponmlkjihgfedcba
```

## Important Constraint

Casting is not allowed for this exercise.

That means you should not convert integer values into runes manually.

Good approach:

- start from rune literal z,
- decrease rune value each loop,
- stop at rune literal a.

This naturally avoids casting.

## New Concept 1: Reverse Loop

In a normal loop, you often do:

- start low,
- go higher,
- stop at upper bound.

In reverse loop, you do the opposite:

- start high,
- go lower,
- stop at lower bound.

For this task:

- start at z,
- decrement by 1,
- continue while value is at least a.

## New Concept 2: Rune Literals

Characters like a and z are runes when written with single quotes:

- a
- z
- newline

This works perfectly with z01.PrintRune, because that function prints a rune.

## Step-By-Step Plan

1. Use package main.
2. Import github.com/01-edu/z01.
3. Create a for loop that starts at z.
4. Loop while current rune is greater than or equal to a.
5. Print each rune with z01.PrintRune.
6. Print newline at the end.

## Final Solution Code

```go
package main

import "github.com/01-edu/z01"

func main() {
	for r := 'z'; r >= 'a'; r-- {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
```

## Why This Is Correct

- It starts from z and moves backward one letter at a time.
- It includes a because condition is greater than or equal.
- It prints newline after letters.
- It uses the allowed function.
- It does not use casting.

## How To Run

From this folder:

```bash
go run .
```

## Common Beginner Mistakes

- using condition r > a, which skips a,
- forgetting newline,
- trying to cast integers to runes,
- using double quotes for characters,
- incrementing instead of decrementing.

## Learning Note

You should now be comfortable with both directions of loops:

- forward order from a to z,
- reverse order from z to a.

That skill will help you in many later Go exercises.