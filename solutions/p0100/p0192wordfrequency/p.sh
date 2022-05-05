#!/usr/bin/env bash

tr '\n' ' ' < words.txt | xargs -n1 | sort | uniq -c | sort -r | awk '{ print $2, $1 }'