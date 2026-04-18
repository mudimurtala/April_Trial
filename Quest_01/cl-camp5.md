# cl-camp5

## Required

- XP: 0.00
- File to submit: lookagain.sh
- Allowed functions: Shell command line tools

## What This Task Is Teaching You

This task continues the file-search theme, but it adds two new ideas:

1. remove part of a filename before printing it,
2. sort results in reverse order.

You are learning how to combine search, text cleanup, and ordering into one short command.

## Task In Simple Words

Create a file named `lookagain.sh`.

When someone runs it, the script should:

- search from the current directory and all subfolders,
- find every file ending with `.sh`,
- print only the file name without `.sh`,
- show the results in descending order.

The example output is:

```text
file3
file2
file1
```

## New Concept 1: Recursive Search

The task says to search from the current directory and its subfolders.

That means the search is recursive.

The command used for that is `find` with `.` as the starting point:

```bash
find .
```

Meaning:

- `.` means current directory,
- `find` will walk into subfolders automatically.

So if there are shell scripts in nested folders, they are included too.

## New Concept 2: Matching File Names

To find shell scripts, match names ending with `.sh`:

```bash
-name '*.sh'
```

Why the quotes matter:

- the `*` is a wildcard,
- quotes stop the shell from expanding it too early,
- `find` should receive the pattern exactly.

## New Concept 3: Showing Only The File Name

The task says the command should only show the name without `.sh`.

That means you must remove the extension before printing.

There are a few ways to do this:

- `sed` can replace the ending `.sh` pattern,
- `cut` can split text by a delimiter,
- `basename` can remove directory parts and, when given a suffix, can also remove `.sh`.

For this task, `basename` is the clearest tool because the input from `find` may include directory paths.

Example:

```bash
basename path/to/file.sh .sh
```

This prints:

```text
file
```

## New Concept 4: Descending Order

The output must be in descending order.

That means reverse alphabetical order for these file names.

Use:

```bash
sort -r
```

What it does:

- `sort` arranges lines in order,
- `-r` reverses that order.

So if the names are `file1`, `file2`, and `file3`, the output becomes:

```text
file3
file2
file1
```

## Why The Hint Mentions Cutting

The hint says a little cutting might be useful.

That is a clue that you need to transform text after searching.

In shell work, "cutting" often means:

- extracting part of a string,
- trimming prefixes or suffixes,
- shaping command output into the exact format required.

In this task, the important concept is not the specific word "cut" itself. The important idea is that you should search first, then trim the `.sh` part, then sort.

## Final Command For The lookagain.sh File

Put this in `lookagain.sh`:

```bash
#!/bin/bash
find . -type f -name '*.sh' -exec basename {} .sh \; | sort -r
```

Why this works:

- `find .` searches the current directory and subfolders,
- `-type f` keeps only files,
- `-name '*.sh'` matches shell scripts,
- `-exec basename {} .sh \;` prints just the name without the `.sh` suffix,
- `sort -r` puts the names in descending order.

this too works:

```bash
#!/bin/bash
find . -type f -name '*.sh' -print0 | xargs -0 -n1 basename -s .sh | sort -r
```
## Step-By-Step To Create And Check It

### 1. Create the file

```bash
touch lookagain.sh
```

### 2. Write the command

```bash
printf '%s\n' '#!/usr/bin/env bash' "find . -type f -name '*.sh' -exec basename {} .sh \\; | sort -r" > lookagain.sh
```

### 3. Make it executable if needed

```bash
chmod +x lookagain.sh
```

### 4. Test it

```bash
./lookagain.sh | cat -e
```

## Common Mistakes To Avoid

- forgetting the quotes around `*.sh`,
- printing full paths instead of only file names,
- forgetting to remove `.sh`,
- using normal sort instead of reverse sort,
- leaving out the shebang.

## Final Reminder

This task teaches a very practical pattern:

- find the files,
- clean the names,
- sort the result,
- print exactly what the checker expects.

If you can do that confidently, you are learning the core rhythm of shell text processing.