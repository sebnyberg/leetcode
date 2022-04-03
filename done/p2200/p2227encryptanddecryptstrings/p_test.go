package p2227encryptanddecryptstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypter(t *testing.T) {
	const typEncrypt = 0
	const typDecrypt = 1
	type any interface{}
	type action struct {
		typ  int
		arg  string
		want any
	}
	for i, tc := range []struct {
		keys       string
		values     []string
		dictionary []string
		actions    []action
	}{
		// {
		// 	keys:       "abcd",
		// 	values:     []string{"ei", "zf", "ei", "am"},
		// 	dictionary: []string{"abcd", "acbd", "adbc", "badc", "dacb", "cadb", "cbda", "abad"},
		// 	actions: []action{
		// 		{typEncrypt, "abcd", "eizfeiam"},
		// 		{typDecrypt, "eizfeiam", 2},
		// 	},
		// },
		{
			keys:       "abcz",
			values:     []string{"aa", "bb", "cc", "zz"},
			dictionary: []string{"aa", "aaa", "aaaa", "aaaaa", "aaaaaa"},
			actions: []action{
				{typDecrypt, "aaaa", 1},
				{typDecrypt, "aa", 1},
				{typDecrypt, "aaaa", 1},
				{typDecrypt, "aaaaaa", 1},
				{typDecrypt, "aaaaaaaaaaaaaa", 1},
				{typDecrypt, "aefagafvabfgshdthn", 0},
			},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			e := Constructor([]byte(tc.keys), tc.values, tc.dictionary)
			for i, a := range tc.actions {
				switch a.typ {
				case typEncrypt:
					require.Equal(t, a.want.(string), e.Encrypt(a.arg), fmt.Sprint(i))
				case typDecrypt:
					require.Equal(t, a.want.(int), e.Decrypt(a.arg), fmt.Sprint(i))
				}
			}
		})
	}
}

type trieNode struct {
	end  bool
	next [26]*trieNode
}

type Encrypter struct {
	charIdx  [26]int
	keys     []byte
	valChars map[string][]int
	vals     [26]string
	valRoot  *trieNode
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	// Create trie from the dictionary.
	// Use traverse trie nodes in parallel based on options provided by values.
	e := Encrypter{
		valChars: make(map[string][]int, 1000),
		valRoot:  new(trieNode),
		keys:     keys,
	}
	for i := range e.charIdx {
		e.charIdx[i] = -1
	}
	for i, k := range keys {
		e.charIdx[k-'a'] = i
	}
	for i, v := range values {
		e.valChars[v] = append(e.valChars[v], i)
		e.vals[i] = v
	}
	for _, d := range dictionary {
		cur := e.valRoot
		for i := range d {
			if cur.next[d[i]-'a'] == nil {
				cur.next[d[i]-'a'] = new(trieNode)
			}
			cur = cur.next[d[i]-'a']
		}
		cur.end = true
	}
	return e
}

func (this *Encrypter) Encrypt(word1 string) string {
	res := make([]byte, 0, len(word1)*2)
	for i := range word1 {
		c := int(word1[i] - 'a')
		idx := this.charIdx[c]
		res = append(res, this.vals[idx][:]...)
	}
	return string(res)
}

func (this *Encrypter) Decrypt(word2 string) int {
	return this.dfs(word2, this.valRoot)
}

func (this *Encrypter) dfs(w string, cur *trieNode) int {
	if len(w) == 0 {
		if cur.end == true {
			return 1
		}
		return 0
	}
	var res int
	for _, ch := range this.valChars[w[:2]] {
		c := this.keys[ch] - 'a'
		if cur.next[c] == nil {
			continue
		}
		res += this.dfs(w[2:], cur.next[c])
	}
	return res
}
