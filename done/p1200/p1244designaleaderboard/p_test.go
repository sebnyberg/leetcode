package p1244designaleaderboard

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type any interface{}

func TestLeaderboard(t *testing.T) {
	l := Constructor()
	type action struct {
		name   string
		inputs []any
		wants  []any
	}
	type testCase struct {
		actions []action
	}
	for i, tc := range []testCase{
		{
			actions: []action{
				{"addScore", []any{1, 73}, nil},
				{"addScore", []any{2, 56}, nil},
				{"addScore", []any{3, 39}, nil},
				{"addScore", []any{4, 51}, nil},
				{"addScore", []any{5, 4}, nil},
				{"top", []any{1}, []any{73}},
				{"reset", []any{1}, nil},
				{"reset", []any{2}, nil},
				{"addScore", []any{2, 51}, nil},
				{"top", []any{3}, []any{141}},
			},
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			l := Constructor()
			for _, act := range tc.actions {
				t.Logf("name:%v\tinputs:%+v\twants:%+v\n", act.name, act.inputs, act.wants)
				switch act.name {
				case "addScore":
					l.AddScore(act.inputs[0].(int), act.inputs[1].(int))
				case "top":
					require.Equal(t, act.wants[0].(int), l.Top(act.inputs[0].(int)))
				case "reset":
					l.Reset(act.inputs[0].(int))
				}
			}
		})
	}
	l.AddScore(1, 73)
}

type Leaderboard struct {
	playersList []*playerScore
	players     map[int]*playerScore
}

type playerScore struct {
	id    int
	score int
}

func Constructor() Leaderboard {
	return Leaderboard{
		playersList: []*playerScore{},
		players:     make(map[int]*playerScore),
	}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
	if _, exists := this.players[playerId]; !exists {
		this.players[playerId] = &playerScore{
			id: playerId,
		}
		this.playersList = append(this.playersList, this.players[playerId])
	}
	this.players[playerId].score += score
}

func (this *Leaderboard) sort() {
	sort.Slice(this.playersList, func(i, j int) bool {
		return this.playersList[i].score >= this.playersList[j].score
	})
}

func (this *Leaderboard) Top(K int) int {
	this.sort()
	var sum int
	for i := 0; i < K; i++ {
		sum += this.playersList[i].score
	}
	return sum
}

func (this *Leaderboard) Reset(playerId int) {
	this.players[playerId].score = 0
}
