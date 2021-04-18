package p0244shortestworddistance

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestDistance(t *testing.T) {
	input := []string{"practice", "makes", "perfect", "coding", "makes"}
	test1 := []string{"coding", "practice"}
	wd := Constructor(input)
	res := wd.Shortest(test1[0], test1[1])
	require.Equal(t, 3, res)
	test2 := []string{"makes", "coding"}
	res = wd.Shortest(test2[0], test2[1])
	require.Equal(t, 1, res)
}

type WordDistance struct {
	wordIndices map[string][]int
	mem         map[[2]string]int
}

func Constructor(wordsDict []string) WordDistance {
	wd := WordDistance{
		wordIndices: make(map[string][]int, len(wordsDict)),
		mem:         make(map[[2]string]int),
	}
	for i, word := range wordsDict {
		wd.wordIndices[word] = append(wd.wordIndices[word], i)
	}
	return wd
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	if word1 < word2 {
		word1, word2 = word2, word1
	}
	k := [2]string{word1, word2}
	if _, exists := this.mem[k]; !exists {
		res := math.MaxInt32
		for _, i := range this.wordIndices[word1] {
			for _, j := range this.wordIndices[word2] {
				res = min(res, abs(i-j))
			}
		}
		this.mem[k] = res
	}
	return this.mem[k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
