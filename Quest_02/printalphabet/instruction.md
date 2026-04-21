# printalphabet

## Required

- Level: 5
- Type: required
- File to submit: `printalphabet/main.go`
- Allowed function: `github.com/01-edu/z01.PrintRune`

## What This Task Is Teaching You

This is a foundational Go exercise.

You are learning how to:

1. write a minimal Go `main` program,
2. print characters one by one using `z01.PrintRune`,
3. use a loop to generate a sequence of letters,
4. end output with a newline (`'\n'`) so it is a proper line.

Even though this looks small, it teaches the exact pattern you will reuse in many early Go tasks.

## Task In Simple Words

Write a program that prints all lowercase Latin letters on one line:

```text
abcdefghijklmnopqrstuvwxyz
```

Then print a newline at the end.

## Important Constraint

You are expected to print runes using:

`z01.PrintRune`

That means you should not use `fmt.Println` for this exercise.

## New Concept 1: Rune vs String

In Go:

- a `rune` is a character (like `'a'`, `'b'`, `'\n'`),
- a string is text (like `"abc"`).

`z01.PrintRune` prints one rune at a time, so this exercise is naturally solved with a loop.

## New Concept 2: ASCII / Unicode Order

Lowercase letters are consecutive in character codes:

- `'a'` comes before `'b'`,
- ...,
- `'z'` is the last lowercase letter.

So a loop from `'a'` to `'z'` prints the alphabet in order.

## Step-By-Step Solution Plan

1. Create `package main`.
2. Import `github.com/01-edu/z01`.
3. In `main`, loop from `rune('a')` to `rune('z')`.
4. Call `z01.PrintRune(letter)` in each loop iteration.
5. Print `z01.PrintRune('\n')` at the end.

## Final Code

```go
package main

import "github.com/01-edu/z01"

func main() {
	for r := 'a'; r <= 'z'; r++ {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
```

## Expected Output

```text
abcdefghijklmnopqrstuvwxyz
```

## How To Run

From inside the exercise folder:

```bash
go run .
```

If your setup runs a specific file:

```bash
go run main.go
```

## Common Mistakes To Avoid

- forgetting the final newline,
- using uppercase letters by mistake,
- using `fmt` instead of `z01.PrintRune`,
- stopping loop before `'z'`,
- starting loop from `'b'` by accident.

## Learning Resource

- `01-edu/z01`: https://github.com/01-edu/z01

This package is used in early exercises to practice character-level output and loop control.
