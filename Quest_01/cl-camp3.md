# cl-camp3

## Required

- XP: 0.00
- File to submit: look
- Allowed functions: Shell command line tools

## What This Task Is Teaching You

This task is about learning the find command properly.

You are learning how to:

1. search from the current directory and all subfolders,
2. match names using wildcard patterns,
3. combine multiple conditions using OR logic,
4. write one correct command inside a file for automated checking.

The main concept here is not just "get output once." It is learning how to ask the filesystem precise questions.

## Task In Simple Words

Create a file named look.

Inside that file, write a command that finds and shows results in the current directory and its subfolders for:

- everything that starts with a,
- all files ending with z,
- all files starting with z and ending with a!.

Hint from the task:
- read the find man page.

## What find Is

find is a Unix command used to search files and directories.

Basic structure:

```bash
find [where-to-search] [conditions]
```

Example:

```bash
find . -name 'a*'
```

Meaning:

- . means start searching in the current directory,
- search is recursive by default (includes subfolders),
- -name 'a*' means match names starting with a.

## Why This Command Matters

Many terminal tasks are really pattern-matching tasks.

If you understand find well, you can:

- locate files quickly in large projects,
- filter by name, type, size, and time,
- combine conditions safely,
- build reliable scripts.

This task is a foundation for all of that.

## Understanding The Patterns In This Task

find with -name uses shell-style wildcard patterns.

### Pattern 1: starts with a

```text
a*
```

- a means the first character is a,
- * means any characters after that (including none).

### Pattern 2: ends with z

```text
*z
```

- * means any starting part,
- z means final character is z.

### Pattern 3: starts with z and ends with a!

```text
z*a!
```

- z means starts with z,
- * means anything in between,
- a! means ends with a!.

Important beginner note:

- put patterns in single quotes so the shell does not expand them before find runs.
- this is especially useful for !, which can behave specially in interactive shells.

## OR Logic In find

You need three alternative matches, so use OR.

In find, OR is written as:

```text
-o
```

So the idea is:

- match condition A OR condition B OR condition C.

To keep logic clear, group conditions using escaped parentheses:

```bash
find . \( condition1 -o condition2 -o condition3 \)
```

Escaping \( and \) is needed so Bash passes parentheses to find instead of treating them as shell syntax.

## About "Everything" vs "Files"

The prompt says:

- everything that starts with a,
- files ending with z,
- files starting with z and ending with a!.

So a careful interpretation is:

- first condition can match both files and directories,
- second and third should be restricted to files.

That is why the final command includes -type f for the second and third conditions.

## Learn From The Manual Page

Run:

```bash
man find
```

Useful searches inside man:

- /-name
- /-type
- /-o

Keys:

- n for next match,
- q to quit.

## Final Command For The look File

Put this exact line in look:

```bash
find . \( -name 'a*' -o -type f -name '*z' -o -type f -name 'z*a!' \)
```

Why this line works:

- starts from current directory and goes through subfolders,
- includes anything whose name starts with a,
- includes files whose names end with z,
- includes files whose names start with z and end with a!.

## Step-By-Step To Create And Verify

### 1. Create the file

```bash
touch look
```

### 2. Add the command into the file

```bash
printf "find . \\( -name 'a*' -o -type f -name '*z' -o -type f -name 'z*a!' \\)\n" > look
```

### 3. Check content

```bash
cat look
```

### 4. Test the command

```bash
bash look
```

## Common Mistakes To Avoid

- forgetting quotes around patterns,
- forgetting escaped parentheses,
- using wrong file name (must be look),
- writing extra lines or unrelated commands.

## Final Reminder

This task trains your ability to search with precision.

Once you can read and combine find conditions confidently, many file-search tasks become easy.