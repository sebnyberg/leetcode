package p1554stringsdifferbyonecharacter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_differByOne(t *testing.T) {
	for _, tc := range []struct {
		dict []string
		want bool
	}{
		{[]string{"moihmlbfdbgdokdknaegcojfbgfhdmbbagk", "moihmibfdbgdokdknjofdkjlfbjhkifgggl", "moicmibfdbgdokdknniijhjnieelambelcd", "moihmibfdbgdokdknlaghflnibmbficaiio", "moihmibfdbgdokdknlehfbjibcicokfblca", "modhmibfdbgdokdknjihgfejeohjkccogcl", "moihmiffdbgdokdknfbedhdoejaofnkbddk", "moihmibfdbgdokdknmdacbcabgaommmaage", "moihmibfdbgdokdknhddnheglfeknalaljg", "moihmibfdbgdokdkncadhfcjfklbibfadhh", "moihmibfdbgdokdkndjoeoaeknbfncbdnbc", "moihmibfdbgdokdknjigmfgbjfndfidofoo", "moihmibfdbgdokdknegbhemjhfaalicgmji", "moihmibfdbgdokdkndejmdgocakfiidkcoa", "moihmibfdbgdokdknbbfcccdifkjgcnndmo", "moihmibfdbgdokdkngighnmaojabljhaaaf", "moihmibfdbgdokdknbhjniooldkdenegdig", "moihmibfdbgdokdkndleffcmmciefjkonbi", "moihmibfdbgdokdkngebdfdkhcoebelcnbg", "mmihmibfdbgdokdkndcgaddaeekmanoccgn", "moihmibfdbgdokdknbabkmoonfejaoaooll", "moihmibfdbgdokdknonadongdkkoibhaogn", "moihmibfdbgdokdknljchjbfcmeeajdaibk", "moihmibfdbgdokdkndibcomalldenlnnjif", "mofhmibfdbgdokdknmgogjmohoofjfcfheh", "moihmibfdbgdokdknmednjgiifonjncdndb", "moihmibfdbgdokdkndficefaecaeafbdnno", "moihmibfdbgdokdknghnhcjkjdciflmnnki", "moihmiffdbgdokdknjdmdlfcjomlibfmdmk", "moihmibfdbgdokdknhoaendcoegbcadecgi", "moihmibfdbgdokdknmncbmnakdlchadjfib", "moihmibfdbgdokdknjgjnbcfodldfoiighd", "moihmibfdbgdokdknoooklmljahdocmogii", "moihmiifdbgdokdklgikhdjfheclhallelc", "moihmibfdbgdofdknlgmojkcjnidgffkglk", "moihmibfdbgdokdknjjejgilcelfcjglcka", "moihmibfdbgdokdkncadhfcjfkliibfadhh", "moihmibfdbgdokdknmlmohehnakhmgmkcfb", "aoihmibfdbgdokdknobcafkenlhnhemmdkc", "moihmibfdbgdokdkncjgajiffhlembllfcm", "moihmibfdbgdokdkngomodaadacglhoedhk", "moihmibfdbgdokdknaenmgfmefkldlhmgma", "moihmibfdbgdokdknhfogboicjdbokoikoa", "moihmibndbgdokdknkndddhnlnjaaomlejg", "moihmibfdbgdokdknnjeldlabkhcoaifhgb", "moihmibfdbgdokdknjogodigbjoebcandll", "moihmibfdbgdokdknhbkmneicinnahjhhfn", "moifmibfdbgdokdknbkjofaofibejhnmcco", "moihmibfdbgdokiknikcogdfimemenojdli"},
			true},
		{[]string{"da", "ae"}, false},
		{[]string{"abcd", "acbd", "aacd"}, true},
		{[]string{"ab", "cd", "yz"}, false},
		{[]string{"abcd", "cccc", "abyd", "abab"}, true},
		{[]string{"mfbdoohablgehdiohdhllhhoodohdnklec", "mfbdoohablgehdiombjjbdlfaeejblknnf", "mfbdoohablgehdiohiokmamdloeajmnhcd", "mfbdoohablgehdiohjlhmdaclacfbjfafa", "mfbdoohablgehdiohnibmonbanccdecjio", "mfbdoohablgehdloheclhdgiekhhgebfka", "mfbdoohablgehdiohacbnlkifleonkchbk", "mfbdoohablgehdiohiiemgjicloiikljag", "mfbdoohablgehdiohhekjcclhnhfehhakg", "mfbdoohablgehdiohkkehgfajgdodbobkk", "mfbdoohablgehdiohhbmklnblkbefmnija", "mfbdoohablgehdiohcfgeeeigbjhkemhfb", "mcbdoohabfgehdioholeoonogjgffnkbcb", "mlbdoohablgehdiohbglfgkhlknakhjhnc", "mfbdoohablgehdiohcbfalkjigjecmohlh", "mfbdoohcblgehdiohlhgncmcnghhkojbfj", "mfbdobhablgehdiohhgidmdgihgbioibhf", "mebdoohablgehdiohlockkkgijomcmbmfj", "mfbdoohablgehdiohmjhdgahlbelglhidh", "mfbdhohablgehdiohlgcnjgbhnaijjhklc", "mfbdoohablgehdiohdknmoolfdaokklmbm", "mfbdoohablgehdiohkcnfmjncllfnhakae", "mfbdoohablgehdiohhdocbblhgakmioijm", "mfbdoohablgehdiobhcmbdmhedcbccgnce", "mfbdoohablgehdiohnekbenflgknalgidn", "mfbdoohablgehdiohmkecdlnoblhoednob", "mfbdoohablgehdiohkbogbjlljeffolgad", "mfbdoohablgehdiohdfllhhoodohdnklec", "mfbdoohablgehdiohadmoljmkgaahoakeo", "mfbdoohablgehdiohnlmgjonokaknoboai", "mfbdoohablgehdiohihcciaifdhkbngkhd", "mfbdoohablgehdiohonbkkhbfblaggdfoa"},
			true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.dict), func(t *testing.T) {
			require.Equal(t, tc.want, differByOne(tc.dict))
		})
	}
}

func differByOne(dict []string) bool {
	// Idea: compute a hash of each element in the dict, then remove one element
	// at a time from the hash. Check if the one-element-removed hashes collide
	// with any previous hashes. If they do, there is a match.
	// mod := 1<<31 - 1
	mod := 1_000_000_007
	// mod := 16_777_619_341_237
	base := 29
	// Create hash of each word
	hashes := make([]int, len(dict))
	for i, w := range dict {
		for _, ch := range w {
			hashes[i] = (hashes[i]*base + int(ch-'a')) % mod
		}
	}

	n := len(dict[0])
	pow := 1
	for j := n - 1; j >= 0; j-- {
		seen := make(map[int][]int)
		for i, w := range dict {
			// Remove current character from hash and check if it has been seen before
			h := (mod + hashes[i] - pow*int(w[j]-'a')%mod) % mod
			for _, otherIdx := range seen[h] {
				if w[:j] == dict[otherIdx][:j] &&
					w[j+1:] == dict[otherIdx][j+1:] {
					return true
				}
			}
			seen[h] = append(seen[h], i)
		}
		pow = (pow * base % mod)
	}
	return false
}
