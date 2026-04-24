# PrintCombN: Combinations of n Different Digits

## Goal

Write a function:

```go
func PrintCombN(n int)
```

It must print all combinations of `n` different digits (`0` to `9`) in ascending order.

`n` is guaranteed to be in this range:
- `0 < n < 10`

## What "Ascending Combinations" Means

Each printed group must have digits in strictly increasing order.

For example, with `n = 3`:
- `012` is valid
- `135` is valid
- `789` is valid
- `001` is invalid (repeated digit)
- `210` is invalid (not increasing)

Also, the full list of combinations must be in ascending order, separated by `, `.

## Output Format

- Combinations are on one line.
- Each combination is separated by comma + space.
- No trailing comma after the last combination.
- End with a newline.

Examples:
- `n = 1` -> `0, 1, 2, 3, 4, 5, 6, 7, 8, 9`
- `n = 3` -> starts at `012` and ends at `789`

## Efficient Strategy (Important)

Because this exercise warns about efficiency, avoid brute-force checking all numbers.

Use a direct combination-building approach:

1. Create a slice of size `n`.
2. Start with the smallest combination:
   - for `n=3`, start as `[0, 1, 2]`
3. Print it.
4. Compute the next valid combination in lexicographic order.
5. Repeat until you reach the last one.

This is efficient because you only generate valid combinations.

## How to Detect the Last Combination

For size `n`, the last combination is:
- `[10-n, 11-n, ..., 9]`

Example:
- if `n=3`, last is `[7, 8, 9]`
- if `n=9`, last is `[1, 2, 3, 4, 5, 6, 7, 8, 9]`

## Single-Function Checker-Friendly Implementation

The version below keeps everything inside `PrintCombN()`.

```go
package main

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	digits := make([]int, n)
	for i := 0; i < n; i++ {
		digits[i] = i
	}

	for {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(digits[i]) + '0')
		}

		last := true
		for i := 0; i < n; i++ {
			if digits[i] != 10-n+i {
				last = false
				break
			}
		}

		if last {
			break
		}

		z01.PrintRune(',')
		z01.PrintRune(' ')

		i := n - 1
		for i >= 0 && digits[i] == 10-n+i {
			i--
		}

		digits[i]++
		for j := i + 1; j < n; j++ {
			digits[j] = digits[j-1] + 1
		}
	}

	z01.PrintRune('\n')
}
```

## Why This Works

- Digits are always strictly increasing.
- You never produce invalid combinations.
- You never repeat combinations.
- You stop exactly at the final valid combination.

## Common Mistakes

1. Printing duplicates like `112`.
2. Printing non-increasing groups like `321`.
3. Leaving a trailing `, ` after the last combination.
4. Using very slow brute-force methods that can timeout.

## Quick Recap

- Build combinations directly, not by filtering random values.
- Keep digits increasing at all times.
- Print separator only between combinations.
- Finish the line with a newline.
