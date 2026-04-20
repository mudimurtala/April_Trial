#!/usr/bin/env bash

SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)

if [ -d "mystery" ]; then
	MYSTERY_DIR="mystery"
elif [ -d "../the-final-cl-test/mystery" ]; then
	MYSTERY_DIR="../the-final-cl-test/mystery"
elif [ -d "$SCRIPT_DIR/../the-final-cl-test/mystery" ]; then
	MYSTERY_DIR="$SCRIPT_DIR/../the-final-cl-test/mystery"
else
	echo "Error: could not find mystery directory." >&2
	exit 1
fi

export KEY_INTERVIEW=$(head -n 179 "$MYSTERY_DIR/streets/Buckingham_Place" | tail -n 1 | grep -o '[0-9]\+')
printf '%s\n' "$KEY_INTERVIEW"
cat "$MYSTERY_DIR/interviews/interview-$KEY_INTERVIEW"
printf '%s\n' "$MAIN_SUSPECT"
