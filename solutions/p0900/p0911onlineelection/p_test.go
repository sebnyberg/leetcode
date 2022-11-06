package p0911onlineelection

import "sort"

type TopVotedCandidate struct {
	times   []int
	leaders []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(persons)
	m := make(map[int]int)
	leaders := make([]int, 0, n)
	var maxVotes int
	var leader int
	for _, p := range persons {
		m[p]++
		if m[p] >= maxVotes {
			maxVotes = m[p]
			leader = p
		}
		leaders = append(leaders, leader)
	}
	return TopVotedCandidate{
		times:   times,
		leaders: leaders,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	i := sort.SearchInts(this.times, t+1) - 1
	return this.leaders[i]
}
