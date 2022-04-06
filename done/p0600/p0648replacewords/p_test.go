package p0648replacewords

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_replaceWords(t *testing.T) {
	for _, tc := range []struct {
		dictionary []string
		sentence   string
		want       string
	}{
		{
			[]string{"e", "k", "c", "harqp", "h", "gsafc", "vn", "lqp", "soy", "mr", "x", "iitgm", "sb", "oo", "spj", "gwmly", "iu", "z", "f", "ha", "vds", "v", "vpx", "fir", "t", "xo", "apifm", "tlznm", "kkv", "nxyud", "j", "qp", "omn", "zoxp", "mutu", "i", "nxth", "dwuer", "sadl", "pv", "w", "mding", "mubem", "xsmwc", "vl", "farov", "twfmq", "ljhmr", "q", "bbzs", "kd", "kwc", "a", "buq", "sm", "yi", "nypa", "xwz", "si", "amqx", "iy", "eb", "qvgt", "twy", "rf", "dc", "utt", "mxjfu", "hm", "trz", "lzh", "lref", "qbx", "fmemr", "gil", "go", "qggh", "uud", "trnhf", "gels", "dfdq", "qzkx", "qxw"},
			"ikkbp miszkays wqjferqoxjwvbieyk gvcfldkiavww vhokchxz dvypwyb bxahfzcfanteibiltins ueebf lqhflvwxksi dco kddxmckhvqifbuzkhstp wc ytzzlm gximjuhzfdjuamhsu gdkbmhpnvy ifvifheoxqlbosfww mengfdydekwttkhbzenk wjhmmyltmeufqvcpcxg hthcuovils ldipovluo aiprogn nusquzpmnogtjkklfhta klxvvlvyh nxzgnrveghc mpppfhzjkbucv cqcft uwmahhqradjtf iaaasabqqzmbcig zcpvpyypsmodtoiif qjuiqtfhzcpnmtk yzfragcextvx ivnvgkaqs iplazv jurtsyh gzixfeugj rnukjgtjpim hscyhgoru aledyrmzwhsz xbahcwfwm hzd ygelddphxnbh rvjxtlqfnlmwdoezh zawfkko iwhkcddxgpqtdrjrcv bbfj mhs nenrqfkbf spfpazr wrkjiwyf cw dtd cqibzmuuhukwylrnld dtaxhddidfwqs bgnnoxgyynol hg dijhrrpnwjlju muzzrrsypzgwvblf zbugltrnyzbg hktdviastoireyiqf qvufxgcixvhrjqtna ipfzhuvgo daee r nlipyfszvxlwqw yoq dewpgtcrzausqwhh qzsaobsghgm ichlpsjlsrwzhbyfhm ksenb bqprarpgnyemzwifqzz oai pnqottd nygesjtlpala qmxixtooxtbrzyorn gyvukjpc s mxhlkdaycskj uvwmerplaibeknltuvd ocnn frotscysdyclrc ckcttaceuuxzcghw pxbd oklwhcppuziixpvihihp",
			"i miszkays w gvcfldkiavww v dvypwyb bxahfzcfanteibiltins ueebf lqhflvwxksi dc k w ytzzlm gximjuhzfdjuamhsu gdkbmhpnvy i mengfdydekwttkhbzenk w h ldipovluo a nusquzpmnogtjkklfhta k nxzgnrveghc mpppfhzjkbucv c uwmahhqradjtf i z q yzfragcextvx i i j gzixfeugj rnukjgtjpim h a x h ygelddphxnbh rvjxtlqfnlmwdoezh z i bbfj mhs nenrqfkbf spfpazr w c dtd c dtaxhddidfwqs bgnnoxgyynol h dijhrrpnwjlju muzzrrsypzgwvblf z h q i daee r nlipyfszvxlwqw yoq dewpgtcrzausqwhh q i k bqprarpgnyemzwifqzz oai pnqottd nygesjtlpala q gyvukjpc s mxhlkdaycskj uvwmerplaibeknltuvd ocnn f c pxbd oklwhcppuziixpvihihp",
		},
		{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery", "the cat was rat by the bat"},
		{[]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs", "a a b c"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.dictionary), func(t *testing.T) {
			require.Equal(t, tc.want, replaceWords(tc.dictionary, tc.sentence))
		})
	}
}

type trieNode struct {
	next [26]*trieNode
	root bool
	w    string
}

func replaceWords(dictionary []string, sentence string) string {
	root := &trieNode{}
	for _, w := range dictionary {
		cur := root
		for _, ch := range w {
			if cur.next[ch-'a'] == nil {
				cur.next[ch-'a'] = &trieNode{}
			}
			cur = cur.next[ch-'a']
		}
		cur.root = true
		cur.w = w
	}

	var strs []string
	for _, w := range strings.Split(sentence, " ") {
		cur := root
		for _, ch := range w {
			cur = cur.next[ch-'a']
			if cur == nil || cur.root {
				break
			}
		}
		if cur == nil || !cur.root {
			strs = append(strs, w)
		} else if cur.root {
			strs = append(strs, cur.w)
		}
	}
	res := strings.Join(strs, " ")
	return res
}
