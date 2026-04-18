# cl-camp1

## Required

- File to submit: `mastertheLS`
- Allowed functions: Shell command line tools

## What This Task Is Teaching You

This task is mainly about learning the `ls` command deeply.

You are learning how to control command output using flags so you can match an exact requirement.

By the end, you should understand:

1. how to list items in your current directory,
2. how hidden files behave,
3. how to sort by access time,
4. how to format output with commas,
5. how to make directory names end with `/`.

## Original Task In Simple Words

In a file named `mastertheLS`, put one command line that does all of this:

- list files and directories in the current folder,
- ignore hidden files (including `.` and `..`),
- separate output with commas,
- sort by access time with newest first,
- add `/` after directory names.

## Start With Terminal Basics

Before writing the final command, make sure you understand where you are and what is inside that location.

Useful commands:

```bash
pwd
ls
```

- `pwd` prints your current directory path.
- `ls` lists visible (non-hidden) files and directories in that location.

If your output is not what you expect, you are probably in the wrong directory. Use `cd` to move into the correct folder.

## Hidden Files: What They Are

In Linux, hidden files usually start with a dot, for example:

- `.git`
- `.env`
- `.bashrc`

Also:

- `.` means current directory
- `..` means parent directory

Important beginner point:

- `ls` (without `-a`) does not show hidden files.
- `ls -a` shows hidden files, including `.` and `..`.

Because this task says ignore hidden files, plain `ls` behavior is what you want.

## How To Learn A Command With `man`

The hint says: Read the man.

`man` means manual page. It is built-in documentation for terminal commands.

For this task:

```bash
man ls
```

How to use the man page:

- type `/comma` and press Enter to search for the word `comma`
- type `/access` to search for access-time options
- press `n` for next match
- press `q` to quit

This is a key skill in terminal work: when you do not remember a flag, check `man`.

## The `ls` Flags You Need Here

These are the important flags for this challenge:

- `-1`
	Force one entry per line (this makes joining with commas reliable).

- `-t`
  Sort by time, newest first.

- `-u`
  Use access time when sorting with `-t`.

- `-p`
  Add `/` to directory names.

For comma-only output with no spaces, you can combine `ls` with `paste -sd, -`.

## Requirement-By-Requirement Mapping

Here is exactly how each requirement is met:

1. list files and directories of current directory:
	`ls` does this.

2. ignore hidden files:
	do not use `-a`; default `ls` already ignores hidden files.

3. separate results with commas only (no spaces):
	use `paste -sd, -` to join lines with commas.

4. order by access time, newest first:
	use `-t -u` (or `-tu`).

5. directories end with `/`:
	use `-p`.

## About The Wording: "ascending" vs "newest first"

The sentence says "ascending order of access time (the newest first)."

These ideas can conflict in strict math wording, but for this task you should follow the explicit expected behavior: newest first.

`ls -tu` gives newest-accessed entries first.

## Do You Need A Loop?

For the official task, no loop is needed.

A short `ls` pipeline is enough and is the cleanest solution.

Still, for learning, loops and operators are useful in shell scripting:

- `for item in *; do ...; done` loops through visible items.
- `[ -d "$item" ]` checks if an item is a directory.
- `&&` runs the next command only if previous command succeeded.
- `||` runs the next command if previous command failed.

You do not need these for this challenge, but you will use them in later tasks.

## What To Put In `mastertheLS`

Put exactly this command line in the `mastertheLS` file:

```bash
ls -1tup | paste -sd, -
```

This command:

- lists current directory entries one per line,
- keeps hidden files out (default behavior),
- joins entries with commas and no spaces,
- sorts by access time, newest first,
- appends `/` to directories.

You can also use this as well:

```bash
ls -utp  --indicator-style=slash | paste -sd, -
```

## Quick Self-Check

Before submitting, run:

```bash
bash mastertheLS
```

or, if the file is executable:

```bash
./mastertheLS
```

Then verify:

- output items are comma-separated,
- directory names end with `/`,
- hidden files are not shown.