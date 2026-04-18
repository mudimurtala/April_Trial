# cl-camp2

## Required

- File to submit: `r`
- Allowed functions: Shell command line tools

## What This Task Is Teaching You

This is a small but important terminal exercise.

You are learning the difference between:

- what is inside a file,
- and what command output should look like.

The goal is to create a file with one correct line, not to write a long script.

## Task In Simple Words

Create a file named `r`.

When someone runs:

```bash
cat -e r
```

the output must be:

```text
R$
```

That means the file contains exactly one line with the letter `R`, followed by a newline.

## What A "Line" Means

A line is text that ends with a newline character (`\n`).

So for this challenge, the file content should be:

- character: `R`
- then newline: `\n`

No extra spaces, no extra letters, no second line.

## Why `cat -e` Is Used

`cat` prints file content.

`cat -e` also shows the end of each line using `$`.

So:

- `R$` means the line contains `R` and then ends correctly.
- If you saw `R $` or `R  $`, that would mean there are extra spaces.
- If you saw two `$` markers on separate lines, that would mean extra lines exist.

## Beginner Step-By-Step

### 1. Make sure you are in the correct directory

```bash
pwd
ls
```

### 2. Create the file with exactly one line

Best safe command:

```bash
printf 'R\n' > r
```

Why `printf` is good here:

- it gives precise control,
- it avoids accidental extra spaces,
- it writes exactly what you ask for.

### 3. Verify the result

```bash
cat -e r
```

Expected:

```text
R$
```

## Common Beginner Mistakes

- Adding a space after `R` (wrong output like `R $`).
- Adding multiple lines.
- Naming the file incorrectly (`R` instead of `r`).
- Forgetting that Linux filenames are case-sensitive.

## What To Put In File `r`

The final file content should be exactly:

```text
R
```

with a newline at the end of the line.

## Final Reminder

This exercise trains attention to output format.

Even for simple tasks, exact formatting matters in automated checks.