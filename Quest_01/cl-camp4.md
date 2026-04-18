# cl-camp4

## Required

- XP: 0.00
- File to submit: myfamily.sh
- Allowed functions: curl, jq, and other shell tools

## What This Task Is Teaching You

This task introduces an important new concept: dynamic input through environment variables.

You are learning how to:

1. read a value from an environment variable (`HERO_ID`),
2. use that value to select one record from JSON,
3. extract the `relatives` field from `connections`,
4. print plain text output without JSON quotes.

This is a real scripting pattern: "read parameter from environment, fetch data, filter result, print exact field."

## Task In Simple Words

Create a file named `myfamily.sh`.

When you run the script, it should:

- use the hero id from environment variable `HERO_ID`,
- download the superhero JSON file,
- find the hero with that id,
- print the hero's family information from `connections.relatives`,
- print it without quotes.

Data source:

```text
https://acad.learn2earn.ng/assets/superhero/all.json
```

## New Concept 1: Environment Variables

An environment variable is a named value provided to commands and scripts.

Example:

```bash
export HERO_ID=1
```

What this does:

- creates a variable named `HERO_ID`,
- makes it available to commands you run after that in the same shell session.

Inside a script, you read it as:

```bash
$HERO_ID
```

Why this matters:

- you can run the same script for different heroes,
- you do not hardcode id values inside the script.

## New Concept 2: Raw Output vs JSON Output

By default, `jq` prints strings as JSON values, which includes quotes.

Example default output:

```text
"Marlo Chandler-Jones (wife); ..."
```

But the task says quotes must be removed.

Use:

```bash
jq -r
```

`-r` means raw output. For strings, it prints plain text without surrounding quotes.

## New Concept 3: Passing Shell Variables Into jq Safely

Do not directly inject shell variables inside jq code text.

Better approach:

```bash
jq --arg id "$HERO_ID" '...'
```

Then use `($id | tonumber)` in jq when matching numeric `id` values.

Why:

- `HERO_ID` comes as text from the shell,
- JSON `id` is a number,
- `tonumber` converts text to number for reliable comparison.

## JSON Path You Need

The required field is:

```text
.connections.relatives
```

So your script logic is:

1. loop records (`.[]`),
2. select the one with matching id,
3. print `.connections.relatives`.

## Build The Command Step By Step

### 1. Fetch JSON

```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json
```

### 2. Filter by id from environment variable

```bash
jq --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber))'
```

### 3. Print relatives only, without quotes

```bash
jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives'
```

## Final Script Content

Put this in `myfamily.sh`:

```bash
#!/bin/bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json \
	| jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives'
```

or this:

```bash
#!/bin/bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq ".[] | select(.id == $HERO_ID) | .connections.relatives" | sed 's/"//g'
```
This is enough to satisfy the task requirement.

## How To Run It

```bash
export HERO_ID=1
chmod +x myfamily.sh
./myfamily.sh
```

or:

```bash
export HERO_ID=1
bash myfamily.sh
```

## Common Mistakes To Avoid

- forgetting to export `HERO_ID`,
- using `jq` without `-r` (output keeps quotes),
- comparing string id to numeric id without conversion,
- printing the whole hero object instead of only `connections.relatives`.

## Final Reminder

The big learning outcome here is parameterized scripting:

- your script reads input (`HERO_ID`) from environment,
- runs the same logic for any id,
- prints exactly one clean field.

That pattern appears in many real shell automation tasks.