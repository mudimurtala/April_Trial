# explain

## Required

- XP: 0.00
- File to submit: explain.sh
- Allowed functions: Shell command line tools

## What This Task Is Teaching You

This task is a summary exercise.

You are not just finding one value anymore. You are reconstructing an investigation trail from several clues and turning that trail into a short, ordered report.

You are learning how to:

1. identify a witness from a clue and an interview file,
2. extract a specific interview number,
3. match a suspect to a car description,
4. compare people across several membership lists,
5. present the final results in the exact order requested.

This is the kind of shell work that combines searching, filtering, and reasoning across multiple files.

## Task In Simple Words

Create a file named `explain.sh`.

When it runs, it must print six lines in this order:

1. the first and last name of the key witness,
2. the interview number of that witness,
3. the colour and make of the main suspect's car,
4. the first of the three other main suspects who were not arrested,
5. the second of those suspects,
6. the third of those suspects.

The output must match the checker exactly.

## New Concept 1: Reading A Trail Across Multiple Files

This exercise is not about one single file.

It asks you to connect clues from:

- the crime scene report,
- the people list,
- the street interview files,
- the vehicles file,
- the memberships lists.

That means the important skill is not just `grep` or `head` by itself. The real skill is combining the output of one clue with the next clue.

## New Concept 2: The Key Witness

The witness is identified by matching the coffee shop clue with the interview and the people file.

The relevant witness is:

```text
Annabel Church
```

The matching interview is:

```text
699607
```

This is the line number or interview reference used to follow the witness trail.

## New Concept 3: The Main Suspect's Car

The car clue comes from the witness's description.

The vehicle that matches the main suspect trail is a:

```text
Blue Honda
```

Important beginner idea:

- the colour comes first,
- the make comes second,
- the final output should keep that order exactly.

## New Concept 4: Matching People Across Membership Lists

The wallet clue gave four membership cards:

- AAA,
- Delta SkyMiles,
- the local library,
- the Museum of Bash History.

The useful technique here is to look for names that appear in all four lists.

Once you do that, the names you keep are the ones that matter for the suspect list.

For this puzzle, the three other main suspects not arrested are:

- Mike Bostock,
- Brian Boyer,
- Matt Waite.

They are shown in alphabetical order of last name.

## Why The Order Matters

The task is strict about ordering.

The three suspect names must be sorted by surname:

1. Bostock
2. Boyer
3. Waite

That is why the output order is:

```text
Mike Bostock
Brian Boyer
Matt Waite
```

## Final Script Content

Put this in `explain.sh`:

```bash
#!/usr/bin/env bash
printf '%s\n' \
	'Annabel Church' \
	'699607' \
	'Blue Honda' \
	'Mike Bostock' \
	'Brian Boyer' \
	'Matt Waite'
```

This prints exactly the six required lines.

## How To Test It

```bash
bash explain.sh | cat -e
```

Expected shape:

```text
Annabel Church$
699607$
Blue Honda$
Mike Bostock$
Brian Boyer$
Matt Waite$
```

## Common Mistakes To Avoid

- mixing up the witness name and the suspect name,
- printing the suspect names in the wrong order,
- forgetting that the output must be exactly six lines,
- adding extra text or labels,
- swapping colour and make.

## Final Reminder

This task is really about explanation and evidence trail.

Once you can connect the clue from the crime scene to the witness, then to the car, then to the membership lists, you are doing real investigative shell work.