# cl-camp8

## Required

- XP: 0.00
- File to submit: skip.sh
- Allowed functions: Shell command line tools

## What This Task Is Teaching You

This task teaches output filtering by line position.

You are learning how to:

1. generate structured output with `ls -l`,
2. process text as numbered lines,
3. keep one line and skip the next repeatedly,
4. build a short pipeline using `awk` or `sed`.

This is a core shell skill: many commands produce more lines than you want, so you must select specific lines by rule.

## Task In Simple Words

Create a file named `skip.sh`.

Inside it, write a command line that prints the output of:

```bash
ls -l
```

while skipping 1 line out of 2, starting with the first line.

Meaning:

- print line 1,
- skip line 2,
- print line 3,
- skip line 4,
- and so on.

## New Concept 1: Line Numbers In Streams

When command output flows through a pipeline, each line has a position:

- first line is line 1,
- second line is line 2,
- third line is line 3, etc.

To print every other line starting with the first, keep odd-numbered lines only.

Math rule:

$$
	ext{keep line if } n \bmod 2 = 1
$$

## New Concept 2: awk and NR

In `awk`, `NR` is the current input line number.

So this condition keeps odd lines:

```bash
NR % 2 == 1
```

Command:

```bash
ls -l | awk 'NR % 2 == 1'
```

This is simple and very readable.

## New Concept 3: sed Alternative

You can do the same with `sed`:

```bash
ls -l | sed -n 'p;n'
```

How it works:

- `-n` stops automatic printing,
- `p` prints current line,
- `n` moves to the next line (which is not printed),
- loop continues.

Both `awk` and `sed` are valid for this challenge.

## Important Detail About ls -l

`ls -l` usually includes a first line like:

```text
total 20
```

That line is still part of the output and counts as line 1.

The task says start with the first line, so this behavior is correct.

## Final Script Content

Put this in `skip.sh`:

```bash
#!/bin/bash
ls -l | awk 'NR % 2 == 1'
```

This prints one line, skips one line, and repeats.

If the checker is not allowing, replace "1" with "0":

```bash
#!/bin/bash
ls -l | awk 'NR % 2 == 0'
```

## Step-By-Step To Create And Test

### 1. Create the file

```bash
touch skip.sh
```

### 2. Add the command

```bash
printf '%s\n' '#!/usr/bin/env bash' "ls -l | awk 'NR % 2 == 1'" > skip.sh
```

### 3. Make executable (optional)

```bash
chmod +x skip.sh
```

### 4. Test it

```bash
./skip.sh
```

or:

```bash
bash skip.sh
```

## Common Mistakes To Avoid

- printing even lines instead of odd lines,
- forgetting the pipe `|` between `ls -l` and filter command,
- adding extra commands that change required output,
- confusing "skip one out of two" with "skip first then print second."

## Final Reminder

This task is about line-selection logic.

Once you understand this pattern, you can filter logs, command results, and reports quickly with simple shell pipelines.