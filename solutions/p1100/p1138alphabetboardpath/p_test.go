package p1138alphabetboardpath

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_alphabetBoardPath(t *testing.T) {
	for i, tc := range []struct {
		target string
		want   string
	}{
		{"leet", "DDR!UURRR!!DDD!"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, alphabetBoardPath(tc.target))
		})
	}
}

func alphabetBoardPath(target string) string {
	pos := [26][2]int{}
	for ch := 0; ch < 26; ch++ {
		pos[ch] = [2]int{
			ch / 5,
			ch % 5,
		}
	}
	var i, j int
	var res []byte
	for _, ch := range target {
		i2 := int((ch - 'a') / 5)
		j2 := int((ch - 'a') % 5)
		for i != i2 || j != j2 {
			if i != i2 {
				d := i2 - i
				if d < 0 {
					i--
					res = append(res, 'U')
				} else if j == 0 || i != 4 {
					i++
					res = append(res, 'D')
				}
			}
			if j != j2 {
				d := j2 - j
				if d < 0 {
					j--
					res = append(res, 'L')
				} else {
					j++
					res = append(res, 'R')
				}
			}
		}
		res = append(res, '!')
	}
	return string(res)
}
