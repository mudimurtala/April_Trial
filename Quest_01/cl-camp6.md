# cl-camp6

## Required

- XP: 0.00
- File to submit: countfiles.sh
- Allowed functions: Shell command line tools

## What This Task Is Teaching You

This task teaches you how to count filesystem entries precisely.

You are learning how to:

1. search through the current directory and all subfolders,
2. count only regular files and directories,
3. include the current directory itself in the result,
4. print only the final number.

This is a useful shell skill because many real tasks are not about listing items, but about counting them correctly.

## Task In Simple Words

Create a file named `countfiles.sh`.

When you run it, the script must print the number of:

- regular files,
- folders,

inside the current directory and all of its subfolders.

Important detail:

- the current directory must be included in the count.

The output must be only the number, for example:

```text
12
```

## New Concept 1: What Counts As A Regular File

A regular file is a normal file that stores data, like:

- text files,
- shell scripts,
- source code files,
- configuration files.

It is not a directory, not a link, and not a device file.

In `find`, a regular file is matched with:

```bash
-type f
```

## New Concept 2: What Counts As A Directory

A directory is a folder.

In `find`, a directory is matched with:

```bash
-type d
```

The current folder `.` is also a directory, so it will be counted when you start the search there.

## New Concept 3: Combining Conditions

You need to count both files and directories.

That means you need OR logic:

```bash
-type f -o -type d
```

The `-o` means OR.

Because the conditions belong together, it is safer to group them:

```bash
\( -type f -o -type d \)
```

## New Concept 4: Counting Lines With wc

`wc` means word count, but it can count lines too.

The option you need is:

```bash
-l
```

So if `find` prints one matching path per line, `wc -l` counts how many lines came through.

That is why this pattern works well:

```bash
find . \( -type f -o -type d \) | wc -l
```

## Why The Current Directory Is Included

The task says the current directory must be included.

This happens naturally because:

- you start the search at `.`,
- `.` is a directory,
- `-type d` matches it.

So the starting folder is counted together with everything inside it.

## Step-By-Step Way To Solve It

### 1. Search from the current directory

```bash
find .
```

This walks through the current directory and all subfolders.

### 2. Keep only regular files and directories

```bash
find . \( -type f -o -type d \)
```

### 3. Count the results

```bash
find . \( -type f -o -type d \) | wc -l
```

### 4. Make sure the output is only the number

`wc -l` normally prints the count with some spacing, but when used in the script it still satisfies the check because the important part is the count itself.

## Final Script Content

Put this in `countfiles.sh`:

```bash
#!/bin/bash
find . \( -type f -o -type d \) | wc -l
```
or this:

```bash
#!/bin/bash
find . | wc -1
```
## How To Run It

```bash
chmod +x countfiles.sh
./countfiles.sh
```

Or:

```bash
bash countfiles.sh
```

## Common Mistakes To Avoid

- forgetting to include the current directory by starting somewhere else,
- counting only files and forgetting directories,
- forgetting to group the `-type` conditions,
- printing extra text instead of only the number.

## Final Reminder

The main lesson here is counting precisely.

`find` gives you the matching entries, and `wc -l` turns those entries into one final number. That is a very common shell pattern you will use again.