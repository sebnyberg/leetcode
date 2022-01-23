#!/usr/bin/env bash

set -euo pipefail

go test -run=None -benchtime=1000x -bench=MinimumOperations -benchmem -cpuprofile=cpu.pprof -memprofile=mem.pprof -trace=trace.out ./done/p2100/p2123minimumoperationstoremovadjacentonesinamatrix/