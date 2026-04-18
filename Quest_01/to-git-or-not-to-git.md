$ to-git-or-not-to-git

## Required

- File to submit: `to-git-or-not-to-git.sh`
- Allowed functions: `curl`, `jq`

## What This Task Is Teaching You

This task teaches you how to read a JSON file from the internet and print only the exact information you need.

You are not writing a large program here. You are writing a tiny shell script that:

1. downloads a JSON file from a URL,
2. finds the superhero with `id` `170`,
3. prints the superhero’s `name`, `power`, and `gender` on separate lines.

That is the full goal of the exercise.

## What The Instructions Mean

The prompt says to write the command in `to-git-or-not-to-git.sh`.

That means the file should contain the shell command that solves the task. When the file is run with Bash, it should print the three values shown in the example output.

The expected output is:

```text
Chameleon
28
Male
```

So the script must find the superhero with `id` `170` and extract:

- `name` from the top level of the JSON object,
- `power` from the `powerstats` section,
- `gender` from the `appearance` section.

## What A JSON File Is

JSON means JavaScript Object Notation.

For a beginner, the easiest way to think about JSON is this:

- it is a text format for structured data,
- it can store many records in a list,
- each record can have named fields inside it.

In this task, each superhero is one record in a big JSON array. Each record has fields like `id`, `name`, `powerstats`, and `appearance`.

## What `curl` Does

`curl` is a command-line tool used to fetch data from a URL.

The JSON file for this task is located at:

```text
https://acad.learn2earn.ng/assets/superhero/all.json
```

When you run `curl` on that URL, it downloads the JSON content and prints it to the terminal.

Useful `curl` flag:

- `-s` means silent mode. It hides extra progress output, which is helpful when you only want clean script output.

## What `jq` Does

`jq` is a tool for reading and filtering JSON.

If `curl` fetches the data, `jq` is the part that looks inside the JSON and picks out the exact values you need.

For this task, `jq` will:

1. go through the list of superheroes,
2. find the one whose `id` is `170`,
3. print three fields from that superhero.

## What The Pipe Symbol Means

You will usually connect `curl` and `jq` with a pipe:

```bash
curl ... | jq ...
```

The pipe symbol `|` means:

- take the output from the command on the left,
- pass it as input to the command on the right.

So `curl` downloads the JSON, and `jq` reads that JSON directly from the pipe.

## How To Read The JSON Path

The task needs values from different places in the record.

The fields are:

- `name` at the top level,
- `powerstats.power` inside the `powerstats` object,
- `appearance.gender` inside the `appearance` object.

The dot notation is how `jq` walks through JSON objects:

- `.name` means get the `name` field,
- `.powerstats.power` means go into `powerstats`, then get `power`,
- `.appearance.gender` means go into `appearance`, then get `gender`.

## Step-By-Step Way To Solve It

Do not try to write the whole command blindly. Build the idea one part at a time.

### 1. Create the script file

Create a file named `to-git-or-not-to-git.sh`.

Start it with a shebang so Bash knows how to run it:

```bash
#!/usr/bin/env bash
```

What this means:

- `#!` marks the file as a script.
- `/usr/bin/env bash` asks the system to run it with Bash.

### 2. Download the JSON

Use `curl` to fetch the superhero data from the URL.

Example idea:

```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json
```

### 3. Find the superhero with id 170

The JSON file contains many superhero records, so you need to select the one where `id` is `170`.

The matching idea in `jq` looks like this:

```bash
jq '.[] | select(.id == 170)'
```

What this means:

- `.[]` loops through every item in the JSON array,
- `select(.id == 170)` keeps only the item whose `id` is `170`.

### 4. Print the three fields you need

Once the correct superhero is selected, print:

- `.name`
- `.powerstats.power`
- `.appearance.gender`

In `jq`, you can print multiple values by separating them with commas. That prints each value on its own line.

### 5. Put everything together

The final command should:

- fetch the JSON with `curl`,
- pass it into `jq`,
- filter the record with `id` `170`,
- print the name, power, and gender.

## What The Final Script Should Look Like

Your `to-git-or-not-to-git.sh` file can look like this:

```bash
#!/bin/bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'
```

or look like this, you can choose any one:

```bash
#!/bin/bash
URL="https://acad.learn2earn.ng/assets/superhero/all.json"
curl -s "$URL" | jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'
```

When you run it with Bash, it should print:

```text
Chameleon
28
Male
```

## Final Reminder

The important things to understand are:

- `curl` fetches the JSON from the web,
- `jq` filters the JSON data,
- `select(.id == 170)` chooses the correct superhero,
- `.name`, `.powerstats.power`, and `.appearance.gender` extract the values you need.

That is the full meaning of the task and the shape of the solution.
$