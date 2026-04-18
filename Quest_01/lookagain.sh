#!/usr/bin/env bash
find . -type f -name '*.sh' -exec basename {} .sh \; | sort -r
