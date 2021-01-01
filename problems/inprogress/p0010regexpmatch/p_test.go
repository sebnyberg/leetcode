package p0010regexpmatch

// func Test_l0010(t *testing.T) {
// 	tcs := []struct {
// 		in   string
// 		in2  string
// 		want bool
// 	}{
// 		// {"a", "ab*a", false},
// 		{"a", "ab*", true},
// 		// {"aa", "a", false},
// 		// {"aa", "a*", true},
// 		// {"aaa", "a*a", true},
// 		// {"ab", ".*", true},
// 		// {"ab", ".*c", false},
// 		// {"aab", "c*a*b", true},
// 		// {"mississippi", "mis*is*p*.", false},
// 	}
// 	for _, tc := range tcs {
// 		t.Run(fmt.Sprintf("%v/%v", tc.in, tc.in2), func(t *testing.T) {
// 			require.Equal(t, tc.want, isMatch(tc.in, tc.in2))
// 		})
// 	}
// }

// func isMatch(s string, p string) bool {
// 	sx := 0
// 	px := 0
// 	for sx < len(p) || px < len(s) {
// 		if px < len(p) {
// 			c := p[px]
// 			switch c {
// 			case '.':
// 				if sx < len(s) {
// 					px++
// 					sx++
// 					continue
// 				}
// 			case '*':
// 				// Try to match at sx, sx+1 and so on
// 				for ; sx < len(s); sx++ {
// 					if isMatch(p[px+1:], s[sx:]) {
// 						return true
// 					}
// 				}
// 			default: // Regular character
// 				if sx < len(s) && s[sx] == c {
// 					px++
// 					sx++
// 					continue
// 				}
// 			}
// 		}
// 		// Mismatch
// 		return false
// 	}
// 	// The entire pattern and string has been matched
// 	return true
// }
