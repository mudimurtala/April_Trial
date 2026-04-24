# PrintComb2: All Ordered Pairs of Two-Digit Numbers

## What This Exercise Wants

You need to write a function named PrintComb2 that prints:
- all possible combinations of two different two-digit numbers
- in ascending order
- on one single line
- separated by comma and space: ", "

Example beginning:
- 00 01, 00 02, 00 03, ...

Example ending:
- ..., 97 98, 97 99, 98 99

## Important Rules

Each item has this format:
- two digits
- one space
- two digits

So combinations look like:
- 00 01
- 01 23
- 98 99

Rules for valid combinations:
- The two numbers must be different.
- The first number must be smaller than the second.
- Output must be in ascending order.

That means:
- 00 01 is valid
- 01 00 is not valid
- 42 42 is not valid

## Allowed Function

Use:
- github.com/01-edu/z01.PrintRune

You should print characters one by one (runes), not with fmt.Println.

## Core Programming Idea

Represent each two-digit number as an integer:
- first from 0 to 98
- second from first+1 to 99

This automatically guarantees:
- first < second
- no duplicates
- ascending order

Then convert each number into tens and units digits:
- tens: n / 10
- units: n % 10

To print digits as characters, add '0':
- rune(n/10) + '0'
- rune(n%10) + '0'

## Step-by-Step Plan

1. Loop first number from 0 to 98.
2. Loop second number from first+1 to 99.
3. Print first as two digits.
4. Print a space.
5. Print second as two digits.
6. Print comma+space after each pair except the last one (98 99).

## Complete Example

```go
package main

import "github.com/01-edu/z01"

func PrintComb2() {
	for first := 0; first <= 98; first++ {
		for second := first + 1; second <= 99; second++ {
			z01.PrintRune(rune(first/10) + '0')
			z01.PrintRune(rune(first%10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune(rune(second/10) + '0')
			z01.PrintRune(rune(second%10) + '0')

			if !(first == 98 && second == 99) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
```

This version keeps everything inside `PrintComb2()`, which is exactly what the checker expects.

If you want to test locally with `go run .`, you can add a temporary `main()` in your own sandbox, but remove it before final submission if your platform expects only the target function.

## Why This Works

- first starts at 0, then 1, then 2, ... up to 98
- second always starts at first+1
- so second is always strictly greater than first
- therefore duplicates and reversed pairs never appear

## Common Mistakes to Avoid

1. Printing a trailing comma after 98 99.
2. Allowing equal pairs like 11 11.
3. Printing reverse pairs like 23 05.
4. Forgetting leading zero (05 must not become 5).

## Quick Recap

- Use nested loops over integers.
- Convert each number to two printable digits.
- Print "first second" format with one space.
- Separate pairs using ", " except for the final pair.