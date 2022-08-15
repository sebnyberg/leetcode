package p0811subdomainvisitcount

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	counts := make(map[string]int)
	for _, s := range cpdomains {
		parts1 := strings.Split(s, " ")
		count, err := strconv.Atoi(parts1[0])
		if err != nil {
			panic(err)
		}
		domain := parts1[1]
		parts2 := strings.Split(domain, ".")
		if len(parts2) > 3 {
			panic(s)
		}
		counts[domain] += count
		counts[parts2[len(parts2)-1]] += count
		if len(parts2) == 3 {
			counts[strings.Join(parts2[1:], ".")] += count
		}
	}
	res := make([]string, 0, len(counts))
	for domain, c := range counts {
		res = append(res, fmt.Sprintf("%v %v", c, domain))
	}
	return res
}
