#!/usr/bin/env bash

# Count words in first column (N)
# Generate a series i := 1 to N 
# Read field i from each row in file.txt
# Combine rows read from cut into a single row with space as the delimiter (- marks stdin as input)
head -1 file.txt | wc -w | xargs seq | xargs -I {} -n1 sh -c "cut -d ' ' -f{} file.txt | paste -sd ' ' -"