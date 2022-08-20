package p0831maskingpersonalinformation

import (
	"fmt"
	"regexp"
	"strings"
)

var stripS = strings.NewReplacer("(", "", ")", "", " ", "", "+", "", "-", "")
var emailR = regexp.MustCompile("^[a-zA-Z]+@[a-zA-Z]+.[a-zA-Z]+$")

func maskPII(s string) string {
	if emailR.MatchString(s) {
		s = strings.ToLower(s)
		parts := strings.Split(s, "@")
		return fmt.Sprintf("%c*****%c@%s", s[0], parts[0][len(parts[0])-1], parts[1])
	}
	s = stripS.Replace(s)
	var pre string
	if len(s) > 10 {
		pre = fmt.Sprintf("+%s-", strings.Repeat("*", len(s)-10))
	}
	return pre + "***-***-" + s[len(s)-4:]
}
