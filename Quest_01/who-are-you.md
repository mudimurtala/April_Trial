# who-are-you

## Required

- Files to submit: `who-are-you.sh`
- Allowed functions: `curl`, `jq`

## What This Task Is Teaching You

This task is a small beginner exercise in reading data from the internet and printing one value from it.

You are not being asked to write a big program. You are being asked to build a tiny shell script that:

1. downloads a JSON file from a URL,
2. looks for the record with `id` equal to `70`,
3. extracts the `name` field from that record,
4. prints only that result.

If you are new to this, the important idea is simple: the script is just a helper that finds one piece of information inside a larger data file.

## What The Clue Means

The instruction says you woke up with a tag that says `subject Id: 70`.

That clue is telling you which record to search for inside the data file.

In the JSON file for this task, the information is stored as a list of objects, and each object has fields like:

- `id`
- `name`
- `appearance`
- `biography`
- `powerstats`

The record you need is the one where `id` is `70`.

## What A JSON File Is

JSON means JavaScript Object Notation.

That sounds technical, but for this task you only need a beginner-friendly idea:

- A JSON file is a text file that stores structured data.
- It can hold lists of items and named values.
- Each item can have fields like `id` and `name`.

Think of it like a neatly organized table in text form.

When you open the superhero JSON file, you are not looking at plain sentences. You are looking at machine-readable data that a tool like `jq` can search.

## What `curl` Does

`curl` is a command-line tool used to fetch data from a URL.

In this task, `curl` is used to download the JSON file from:

```text
https://acad.learn2earn.ng/assets/superhero/all.json
```

Beginner idea:

- `curl` goes to the web address.
- It brings the file contents back to your terminal.

Useful `curl` flag:

- `-s` means silent mode. It keeps `curl` from showing extra progress output.

That is helpful when you want your script to print only the final answer.

## What `jq` Does

`jq` is a tool for reading and filtering JSON.

If `curl` brings the JSON file to your terminal, `jq` is the tool that helps you pick out the exact piece you want.

Beginner idea:

- `curl` gets the file.
- `jq` searches inside the file.
- `jq` prints the matching value.

For this task, you use `jq` to:

1. look through every item in the JSON array,
2. find the one whose `id` is `70`,
3. print the `name` field from that item.

## What The Pipe Symbol Means

You will usually connect `curl` and `jq` with a pipe:

```bash
curl ... | jq ...
```

The pipe symbol `|` means:

- take the output from the command on the left,
- send it into the command on the right.

So in this task, the JSON downloaded by `curl` becomes the input that `jq` reads.

## Why The Output Has Quotes

The example output shows a quoted name, like this:

```text
"Batman"
```

That is because `jq` prints JSON values in JSON style by default.

If you use the raw output flag `-r`, the quotes would disappear, but this exercise expects the JSON-style output shown in the example.

That means the final answer should look like one quoted name on one line, with nothing else added.

## Step-By-Step Way To Solve It

Do not try to guess the whole command at once. Build it one piece at a time.

### 1. Create the shell script

Create a file named `who-are-you.sh`.

Start the file with a shebang so the system knows how to run it:

```bash
#!/usr/bin/env bash
```

What the shebang means:

- `#!` tells the system this file is a script.
- `/usr/bin/env bash` asks the system to run the script with Bash.

### 2. Fetch the JSON data

Use `curl` to download the file from the URL.

Example idea:

```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json
```

This gives you the full JSON content.

### 3. Filter the record you need

Use `jq` to find the object whose `id` is `70`.

The pattern you want is:

```bash
jq '.[] | select(.id == 70) | .name'
```

What this means:

- `.[]` goes through each item in the JSON array.
- `select(.id == 70)` keeps only the item with `id` equal to `70`.
- `.name` prints the name field from that item.

### 4. Put the commands together

The script should connect the download step and the filtering step with a pipe.

In plain English, the final script should:

- download the JSON,
- search for `id 70`,
- print the matching `name`.

### 5. Test the result

Run the script and check the output.

The output should be exactly one line, similar to this example:

```text
"Batman"
```

The exact name may be different if the data changes, but the format should stay the same.

## What The `who-are-you.sh` File Should Look Like

Your file should be a small Bash script, like this:

```bash
#!/bin/bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name'
```

This is the core idea of the exercise.

If you want to make the script more careful, you can keep the same logic and still ensure the output stays clean and readable.

## Final Reminder

The important part is not just getting the answer. It is understanding the tools:

- `curl` brings data from a URL.
- `jq` reads JSON and extracts the matching value.
- the pipe `|` connects them.
- the script prints the name for the record with `id` `70`.

That is how `who-are-you.sh` should behave.
$