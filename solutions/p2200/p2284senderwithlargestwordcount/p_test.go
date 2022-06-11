package p2284senderwithlargestwordcount

import "strings"

func largestWordCount(messages []string, senders []string) string {
	m := make(map[string]int)
	for i := range messages {
		m[senders[i]] += len(strings.Fields(messages[i]))
	}

	var maxCount int
	var res string
	for k, v := range m {
		if v > maxCount {
			res = k
			maxCount = v
		} else if v == maxCount && k > res {
			res = k
		}
	}
	return res
}
