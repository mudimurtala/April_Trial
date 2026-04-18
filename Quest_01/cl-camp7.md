# cl-camp7

## Required

- XP: 0.00
- File to submit: `"\?$*'ChouMi'*$?\"`
- Allowed functions: Shell command line tools

## What This Task Is Teaching You

This task teaches precision with shell quoting and special characters in filenames.

You are learning how to:

1. create a file whose name contains special characters,
2. avoid shell wildcard expansion,
3. avoid variable expansion by mistake,
4. write exact file content (`01` and nothing else).

This is a very important terminal skill because shells interpret many symbols unless you quote correctly.

## Task In Simple Words

Create a file whose name is exactly:

```text
"\?$*'ChouMi'*$?\"
```

And the file content must be exactly:

```text
01
```

with nothing else added.

## New Concept 1: Why This Filename Is Tricky

The filename includes characters that usually have special meaning in shell:

- `*` wildcard (matches many names),
- `$` variable marker,
- `?` single-character wildcard,
- `"` quote characters,
- `'` single quotes inside the filename.

If you type this name without careful quoting, the shell may transform it before creating the file.

## New Concept 2: Quoting Rules You Need

### Single quotes `'...'`

- protect almost everything literally,
- but cannot directly contain another single quote.

### Double quotes `"..."`

- still allow `$` variable expansion,
- still need escaping for some characters.

For this task, the easiest safe method is:

1. store the target filename in a variable using double quotes and escapes,
2. always reference it as `"$name"`.

## New Concept 3: Writing Exactly `01`

Use `printf` instead of `echo`.

Why:

- `echo` behavior can vary,
- `printf` gives exact output control,
- `printf '01'` writes exactly two characters with no automatic extra newline.

Since the task says "01 and nothing else," `printf '01'` is the safest choice.

## Step-By-Step Safe Solution

### 1. Define the filename exactly

```bash
name="\"\\?\$*'ChouMi'*\$?\\\""
```

### 2. Create the file with exact content

```bash
printf '01' > "$name"
```

### 3. Verify the file exists

```bash
ls | cat -e
```

You should see the exact filename in the list.

### 4. Verify content is exact

```bash
cat "$name" | cat -e
```

Expected content line should show only `01` (and no extra characters).

## Alternative One-Liner

If you prefer one compact command:

```bash
name="\"\\?\$*'ChouMi'*\$?\\\"" && printf '01' > "$name"
```

## Common Mistakes To Avoid

- typing the filename without quotes,
- letting `*` expand into real files,
- letting `$` try to read shell variables,
- using `echo` and accidentally adding extra content,
- creating a similar but not exact filename.

## Final Reminder

This task is about accuracy, not complexity.

In shell, exact quoting is often the difference between a correct solution and a hidden bug.