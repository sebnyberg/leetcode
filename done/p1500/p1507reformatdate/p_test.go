package p1507reformatdate

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reformatDate(t *testing.T) {
	for _, tc := range []struct {
		date string
		want string
	}{
		{"20th Oct 2052", "2052-10-20"},
		{"6th Jun 1933", "1933-06-06"},
		{"26th May 1960", "1960-05-26"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.date), func(t *testing.T) {
			require.Equal(t, tc.want, reformatDate(tc.date))
		})
	}
}

var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

func reformatDate(date string) string {
	// 1st 2nd 3rd 4th 5th 6th 7th 8th 9th 10th ...
	parts := strings.Split(date, " ")
	day, _ := strconv.Atoi(parts[0][:len(parts[0])-2])
	var month int
	for i, mo := range months {
		if mo == parts[1] {
			month = i + 1
		}
	}
	dayStr := strconv.Itoa(day)
	if day < 10 {
		dayStr = "0" + dayStr
	}
	moStr := strconv.Itoa(month)
	if month < 10 {
		moStr = "0" + moStr
	}
	yearStr := parts[2]

	return yearStr + "-" + moStr + "-" + dayStr
}
