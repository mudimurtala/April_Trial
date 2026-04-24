# PrintComb: All Increasing 3-Digit Combinations

## Goal of This Exercise

Write a function named `PrintComb` that prints all valid combinations of 3 different digits in increasing order.

A combination is valid only if:
- The first digit is smaller than the second.
- The second digit is smaller than the third.
- All digits are different.

So:
- `012` is valid.
- `019` is valid.
- `123` is valid.
- `789` is valid.
- `000` is invalid (digits repeat).
- `987` is invalid (not increasing).

All combinations must be printed:
- On one single line
- In ascending order
- Separated by `, ` (comma + space)

## Important Constraint

You are allowed to use:
- `github.com/01-edu/z01.PrintRune`

And this rule appears in the subject:
- `--no-lit=\d{3},`

That means you should not hardcode 3-digit values like `123` followed by a comma as numeric literals. The proper way is to build output character by character with loops and `PrintRune`.

## Core Idea

Use 3 nested loops:
- Outer loop: first digit (`a`)
- Middle loop: second digit (`b`)
- Inner loop: third digit (`c`)

To guarantee increasing digits automatically:
- `b` should start at `a + 1`
- `c` should start at `b + 1`

This enforces:
- `a < b < c`

which means all three digits are automatically different and sorted.

## Step-by-Step Plan

1. Loop `a` from `'0'` to `'7'`
2. Loop `b` from `a + 1` to `'8'`
3. Loop `c` from `b + 1` to `'9'`
4. Print `a`, then `b`, then `c`
5. Print `, ` after each combination except the last one (`789`)

## Why the Last One Matters

You must not print an extra trailing comma at the end.

Correct ending:
- `..., 689, 789`

Incorrect ending:
- `..., 689, 789, `

So add a condition: only print comma+space if the current combination is not `789`.

## Complete Example

```go
package main

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)

				if !(a == '7' && b == '8' && c == '9') {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}

func main() {
	PrintComb()
}
```

## What You Learn Here

- How to generate combinations with nested loops
- How to control ordering with loop starting points
- How to avoid duplicates naturally
- How to handle separator formatting cleanly
- How to print output rune by rune

## Quick Recap

- We need all 3-digit combinations with strictly increasing digits.
- Using `a < b < c` gives correct combinations only.
- Print each combo as characters.
- Separate with `, ` except after `789`.
- Output is on a single line.