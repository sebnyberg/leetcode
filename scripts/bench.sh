#!/usr/bin/env bash

set -eu

rm -rf tmp/old.txt tmp/new.txt tmp/res.txt
[[ ! -e "tmp" ]] && mkdir
touch tmp/res.txt

for i in {1..5} ; do
    go test -run=None -bench=BenchmarkA -benchmem ./tmp | tee tmp/old.txt 
done

benchstat tmp/old.txt
# benchstat tmp/old.txt tmp/new.txt
# rm -rf tmp

# go test -count 1 -run=None -bench=Day19Part2 -benchmem -cpuprofile=cpu.pprof -memprofile=mem.pprof -trace=trace.out ./aoc2021/day19/...