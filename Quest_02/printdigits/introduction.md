# printdigits

## Required

- Level: 5
- Type: required
- File to submit: printdigits/main.go
- Allowed function: github.com/01-edu/z01.PrintRune

## What This Exercise Is Teaching You

This exercise teaches character output and loop basics in Go.

You are learning how to:

1. print digits in ascending order,
2. use a loop to produce a sequence,
3. print one rune at a time with z01.PrintRune,
4. end output with a newline.

This is one of the core beginner patterns in 01-edu Go exercises.

## Task In Simple Words

Write a Go program that prints:

0123456789

on one line, then prints newline.

## Output Rule

A line means the characters must be followed by end-of-line ('\n').

Expected output:

```text
0123456789
```

## New Concept 1: Digits Are Characters Here

In this task, you are printing characters '0' through '9', not doing arithmetic output with fmt.

That means each printed item is a rune:

- '0'
- '1'
- ...
- '9'

## New Concept 2: Sequential Character Range

Character codes for '0' to '9' are consecutive.

So a loop from '0' to '9' prints all decimal digits in order.

## Step-By-Step Plan

1. Use package main.
2. Import github.com/01-edu/z01.
3. Create a loop starting at '0' and ending at '9'.
4. Print each rune using z01.PrintRune.
5. Print newline at the end.

## Final Code

```go
package main

import "github.com/01-edu/z01"

func main() {
	for r := '0'; r <= '9'; r++ {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
```

## Why This Is Correct

- Starts at '0'.
- Increments one step each loop.
- Stops after '9'.
- Prints newline so output is exactly one proper line.

## How To Run

From this folder:

```bash
go run .
```

## Common Beginner Mistakes

- forgetting the final newline,
- using fmt.Println instead of z01.PrintRune,
- starting from 0 instead of '0',
- stopping at '8' by mistake.

## Learning Resource

- 01-edu/z01: https://github.com/01-edu/z01
