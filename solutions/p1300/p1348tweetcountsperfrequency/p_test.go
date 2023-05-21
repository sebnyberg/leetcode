package p1348tweetcountsperfrequency

import (
	"math"
	"sort"
)

type TweetCounts struct {
	tweets map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{
		tweets: make(map[string][]int),
	}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	if _, exists := this.tweets[tweetName]; !exists {
		this.tweets[tweetName] = append(make([]int, 0, 2), -1, math.MaxInt64)
	}
	i := sort.SearchInts(this.tweets[tweetName], time)
	this.tweets[tweetName] = append(this.tweets[tweetName], 0)
	copy(this.tweets[tweetName][i+1:], this.tweets[tweetName][i:])
	this.tweets[tweetName][i] = time
}

var dist = map[string]int{
	"minute": 60,
	"hour":   60 * 60,
	"day":    60 * 60 * 24,
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	if _, exists := this.tweets[tweetName]; !exists {
		return []int{}
	}
	tt := this.tweets[tweetName]
	j := sort.SearchInts(tt, startTime)
	d := dist[freq]
	var res []int
	for i := startTime; i <= endTime; i += d {
		var count int
		for j < len(tt) && tt[j] >= i && tt[j] < min(endTime+1, i+d) {
			count++
			j++
		}
		res = append(res, count)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
