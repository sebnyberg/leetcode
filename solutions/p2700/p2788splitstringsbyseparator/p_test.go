package p2788splitstringsbyseparator

import "strings"

func splitWordsBySeparator(words []string, separator byte) []string {
	var res []string
	for _, w := range words {
		for _, ss := range strings.Split(w, string(separator)) {
			if ss == "" {
				continue
			}
			res = append(res, ss)
		}
	}
	return res
}
