#!/usr/bin/env bash

set -eu

rm -rf tmp
mkdir tmp
touch tmp/res.txt

for i in {1..200} ; do
    # go test -run=None -bench=BenchmarkHashSet -benchmem ./done/p0700/p0705designhashsetotherbefore/... | grep HashSet | tee tmp/old.txt 
    go test -run=None -bench=BenchmarkHashSet -benchmem ./done/p0700/p0705designhashsetotherafter/... | grep HashSet | tee tmp/old.txt 
done

# benchstat tmp/old.txt tmp/new.txt
# rm -rf tmp

# go test -count 1 -run=None -bench=Day19Part2 -benchmem -cpuprofile=cpu.pprof -memprofile=mem.pprof -trace=trace.out ./aoc2021/day19/...