package main

import "github.com/pkg/profile"

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	leastBricks([][]int{{100000000}})
}

func leastBricks(wall [][]int) int {
	n := sum(wall[0]) - 1
	if n == 0 {
		return len(wall)
	}
	edges := make(map[int]int)
	e := edgeFinder{
		buf: make([]int, 0),
	}
	for _, brickRow := range wall {
		for _, edgeIdx := range e.findEdges(brickRow) {
			edges[edgeIdx] = edges[edgeIdx] + 1
		}
	}
	maxVal := 0
	for _, v := range edges {
		maxVal = max(maxVal, v)
	}
	return len(wall) - maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(ns []int) int {
	var res int
	for _, n := range ns {
		res += n
	}
	return res
}

type edgeFinder struct {
	buf []int
}

func (e *edgeFinder) findEdges(brickRow []int) []int {
	var sum int
	e.buf = e.buf[:0]
	for i, brickWidth := range brickRow {
		if i == len(brickRow)-1 {
			break
		}
		sum += brickWidth
		e.buf = append(e.buf, sum-1)
	}
	return e.buf
}
