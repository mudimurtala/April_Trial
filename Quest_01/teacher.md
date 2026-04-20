# teacher

## Required

- XP: 1.65 kB
- File to submit: teacher.sh
- Allowed functions: Shell command line tools

## What This Task Is Teaching You

This task is about automation across changing datasets.

You are learning how to:

1. extract a value from a file using a stable path,
2. store that value in an environment variable,
3. reuse that variable to read another file,
4. print another environment variable already provided by the test system.

This is a key scripting idea: do not hardcode values that change. Compute them from data.

## Task In Simple Words

Write a file named teacher.sh that does these four steps in order:

1. find the key interview number and store only the number in an environment variable,
2. print that environment variable,
3. print the interview content that corresponds to that number,
4. print the content of MAIN_SUSPECT.

Important detail from the prompt:

- the interview number can change,
- but the path used to discover it stays the same across mystery folders.

## New Concept 1: Command Substitution

Command substitution lets you run a command and store its output in a variable.

Syntax:

```bash
var=$(some_command)
```

For this task, that means:

- run a command that returns the interview number,
- store that number into a variable.

## New Concept 2: Environment Variable vs Local Variable

A shell variable is local by default.

An environment variable is exported and available to child processes.

Use:

```bash
export KEY_INTERVIEW=...
```

This satisfies the requirement that the number is isolated into an environment variable.

## New Concept 3: Extracting Only The Number

The line in the street file looks like:

```text
SEE INTERVIEW #699607
```

You must keep only:

```text
699607
```

Use a number-only extractor like:

```bash
grep -o '[0-9]\+'
```

That prints only the numeric part.

## New Concept 4: Reusing The Extracted Number

Once the number is in KEY_INTERVIEW, build the interview path dynamically:

```bash
mystery/interviews/interview-$KEY_INTERVIEW
```

Then print its content with cat.

This keeps the script dynamic even when interview IDs change.

## Why This Specific Path Works

The prompt says the path used to find the interview is stable.

So the same line-slice command can be used each time.

A reliable pattern is:

```bash
head -n 179 mystery/streets/Buckingham_Place | tail -n 1
```

Then extract only the number from that line.

## Final Script Content

Put this in teacher.sh:

```bash
#!/usr/bin/env bash

export KEY_INTERVIEW=$(head -n 179 mystery/streets/Buckingham_Place | tail -n 1 | grep -o '[0-9]\+')
printf '%s\n' "$KEY_INTERVIEW"
cat "mystery/interviews/interview-$KEY_INTERVIEW"
printf '%s\n' "$MAIN_SUSPECT"
```

## How To Test

From the folder that contains mystery:

```bash
export MAIN_SUSPECT="Some Name"
bash teacher.sh
```

You should see, in order:

1. the interview number,
2. the interview text,
3. the MAIN_SUSPECT value.

## Common Mistakes To Avoid

- hardcoding a fixed interview number,
- printing text labels like "Interview:" or "Suspect:" (unless the task asks),
- forgetting to extract digits only,
- forgetting to use the variable when reading interview content,
- forgetting to print MAIN_SUSPECT at the end.

## Final Reminder

This exercise is automation practice.

A good script should survive data changes by recomputing values from stable clues.