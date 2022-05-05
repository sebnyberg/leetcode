package p0355designtwitter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwitter(t *testing.T) {
	twitter := Constructor()
	twitter.PostTweet(1, 5)       // User 1 posts a new tweet (id = 5).
	res := twitter.GetNewsFeed(1) // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
	require.Equal(t, []int{5}, res)
	twitter.Follow(1, 2)         // User 1 follows user 2.
	twitter.PostTweet(2, 6)      // User 2 posts a new tweet (id = 6).
	res = twitter.GetNewsFeed(1) // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	require.Equal(t, []int{6, 5}, res)
	twitter.Unfollow(1, 2)       // User 1 unfollows user 2.
	res = twitter.GetNewsFeed(1) // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
	require.Equal(t, []int{5}, res)
}

type Twitter struct {
	follows map[uint16]map[uint16]struct{}
	tweets  []tweet
}

type tweet struct {
	ID     uint16
	userID uint16
}

func Constructor() Twitter {
	return Twitter{
		tweets:  make([]tweet, 0, 10001),
		follows: make(map[uint16]map[uint16]struct{}, 500),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, exists := this.follows[uint16(userId)]; !exists {
		this.follows[uint16(userId)] = make(map[uint16]struct{}, 10)
		this.follows[uint16(userId)][uint16(userId)] = struct{}{}
	}
	this.tweets = append(this.tweets, tweet{
		ID:     uint16(tweetId),
		userID: uint16(userId),
	})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	// Brute-force solution first: iterate over tweets by id desc, add tweets
	// until reaching 10 in total or end of list of tweets.
	res := make([]int, 0, 10)
	for i := len(this.tweets) - 1; i >= 0; i-- {
		tw := this.tweets[i]
		if _, exists := this.follows[uint16(userId)][uint16(tw.userID)]; exists {
			res = append(res, int(tw.ID))
			if len(res) == 10 {
				return res
			}
		}
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, exists := this.follows[uint16(followerId)]; !exists {
		this.follows[uint16(followerId)] = make(map[uint16]struct{}, 10)
		this.follows[uint16(followerId)][uint16(followerId)] = struct{}{}
	}
	this.follows[uint16(followerId)][uint16(followeeId)] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, exists := this.follows[uint16(followerId)]; !exists {
		this.follows[uint16(followerId)] = make(map[uint16]struct{}, 10)
		this.follows[uint16(followerId)][uint16(followerId)] = struct{}{}
	}
	delete(this.follows[uint16(followerId)], uint16(followeeId))
}
