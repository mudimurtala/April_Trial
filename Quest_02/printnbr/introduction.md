# PrintNbr: Print Any Integer Value

## Exercise Goal

Write a function `PrintNbr(n int)` that prints the integer `n` using only `z01.PrintRune`.

Key requirements:
- It must work for all possible `int` values.
- You cannot convert `int` to `int64`.
- You should produce characters, not use `fmt`.

## Why This Exercise Is Important

This exercise teaches two core ideas:
- How to print a number digit by digit.
- How to handle edge cases safely (especially negative values and MinInt).

## Main Challenges

1. **Negative numbers**
- If `n` is negative, print `'-'` first.

2. **Multiple digits**
- A number like `123` must print as `'1'`, `'2'`, `'3'`.

3. **The MinInt trap**
- The smallest `int` value cannot be directly negated (`-n` may overflow).
- So we avoid turning the full number positive at once.

## Safe Strategy (No int64, Works for MinInt)

Use recursion:

1. Print `'-'` if `n < 0`.
2. If the number has more than one digit (`n <= -10` or `n >= 10`):
   - For positive `n`, recurse with `n / 10`.
   - For negative `n`, recurse with `-(n / 10)`.
3. Print the last digit using `n % 10`.
4. If last digit is negative, make that digit positive before printing.

Why this avoids overflow:
- We never do `-n` for the full number.
- We only negate smaller pieces like `n/10` or `n%10`, which are safe.

## Character Conversion Reminder

Digits are printed as runes:
- `0` becomes `'0'`
- `1` becomes `'1'`
- ...

Formula:
- `z01.PrintRune(rune(digit) + '0')`

## Complete Solution Function

```go
package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}

	if n <= -10 || n >= 10 {
		if n < 0 {
			PrintNbr(-(n / 10))
		} else {
			PrintNbr(n / 10)
		}
	}

	digit := n % 10
	if digit < 0 {
		digit = -digit
	}
	z01.PrintRune(rune(digit) + '0')
}
```

This version keeps everything inside the expected function and handles all `int` values correctly.

## Example

If you call:
- `PrintNbr(-123)` prints `-123`
- `PrintNbr(0)` prints `0`
- `PrintNbr(123)` prints `123`

Combined output: `-1230123`

## Common Mistakes to Avoid

1. Using `fmt.Println` instead of `z01.PrintRune`.
2. Doing `n = -n` directly for all negatives (can break for MinInt).
3. Forgetting to handle `0`.
4. Printing only the last digit and forgetting recursion for earlier digits.

## Quick Recap

- Print sign first if needed.
- Recurse to print higher-order digits.
- Print the final digit.
- Avoid full negation of `n` to stay safe for MinInt.