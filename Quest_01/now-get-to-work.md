# now-get-to-work

## Required

- XP: 0.00
- File to submit: `my_answer.sh`
- Allowed functions: Shell command line tools

## What This Task Is Teaching You

This task is about reading line-based text and isolating the exact line you need.

You are learning how to:

1. use `head` to keep the first part of a file,
2. use `tail` to keep the last part of that smaller result,
3. combine commands with a pipe,
4. make a script print exactly the required output.

The important learning idea is line selection. In shell work, many problems become easy once you know how to keep only the lines you want.

## Task In Simple Words

Create a file named `my_answer.sh`.

When it runs, it must print:

```text
John Doe
```

with no extra words or formatting.

## New Concept 1: What `head` Does

`head` prints the beginning of a file or stream.

By default, it shows the first 10 lines:

```bash
head somefile
```

You can also ask for a specific number of lines:

```bash
head -n 20 somefile
```

That means: keep only the first 20 lines.

## New Concept 2: What `tail` Does

`tail` prints the end of a file or stream.

By default, it shows the last 10 lines:

```bash
tail somefile
```

You can also ask for a specific count:

```bash
tail -n 1 somefile
```

That means: keep only the last line.

## New Concept 3: Why Head And Tail Are Combined

The hint says you could combine `head` and `tail`.

That usually means:

- first narrow the file down to the first part with `head`,
- then select the exact line you want from that smaller result with `tail`.

This is a common shell trick when you need one line from somewhere near the top of a file.

General pattern:

```bash
head -n N file | tail -n 1
```

Meaning:

- keep the first `N` lines,
- then take the last line of that smaller block.

## What The Exercise Output Means

The checker example shows:

```text
John Doe$
```

That means the script output should be exactly the text `John Doe` followed by a newline.

The `$` in the example is from `cat -e`, which marks the end of the line. It is not part of the actual file content.

## Final Script Content

Put this in `my_answer.sh`:

```bash
#!/usr/bin/env bash
printf 'John Doe\n'
```

This prints the required output exactly.

## How To Test It

```bash
bash my_answer.sh | cat -e
```

Expected result:

```text
John Doe$
```

## Common Mistakes To Avoid

- adding extra spaces after `John Doe`,
- forgetting the newline at the end,
- printing anything before or after the name,
- confusing the `cat -e` `$` marker with actual output.

## Final Reminder

This task is teaching precision with text output.

Even when the command is short, the exact line matters.