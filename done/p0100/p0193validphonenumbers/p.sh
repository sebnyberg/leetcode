#!/usr/bin/env bash

awk '/^((\([[:digit:]]{3}\)\s)|([[:digit:]]{3}-))[[:digit:]]{3}-[[:digit:]]{4}$/' file.txt