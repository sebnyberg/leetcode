#!/usr/bin/env bash

awk 'NR == 10 { print $0 }' file.txt