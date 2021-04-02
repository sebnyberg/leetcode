package p0157readncharsgivenread4part2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//
// The read4 API is already defined for you.
//
//     read4 := func(buf4 []byte) int
//
// Below is an example of how the read4 API can be called.
// file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
// buf4 := make([]byte, 4) // Create buffer with enough space to store characters
// read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
// read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
// read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
//

type FakeReader struct {
	contents string
	pos      int
}

func (fr *FakeReader) getRead4() func([]byte) int {
	return func(b []byte) int {
		if fr.pos >= len(fr.contents) {
			return 0
		}
		if len(fr.contents[fr.pos:]) < 4 {
			copy(b, fr.contents[fr.pos:])
			nread := len(fr.contents) - fr.pos
			fr.pos = len(fr.contents)
			return nread
		}
		copy(b, fr.contents[fr.pos:fr.pos+4])
		fr.pos += 4
		return 4
	}
}

func TestSolution(t *testing.T) {
	fr := FakeReader{
		contents: "1234567",
	}
	read := solution(fr.getRead4())
	for _, tc := range []struct {
		ntoread int
		want    string
	}{
		{1, "1"},
		{1, "2"},
		{1, "3"},
		{1, "4"},
		{1, "5"},
		{1, "6"},
		{1, "7"},
	} {
		b := make([]byte, len(fr.contents))
		nread := read(b, tc.ntoread)
		b = b[:nread]
		require.Equal(t, tc.want, string(b))
	}
}

func TestSolution2(t *testing.T) {
	fr := FakeReader{
		contents: "abc",
	}
	read := solution(fr.getRead4())
	for _, tc := range []struct {
		ntoread int
		want    string
	}{
		{4, "abc"},
		{1, ""},
	} {
		b := make([]byte, len(fr.contents))
		nread := read(b, tc.ntoread)
		b = b[:nread]
		require.Equal(t, tc.want, string(b))
	}
}

func TestSolution3(t *testing.T) {
	fr := FakeReader{
		contents: "abc",
	}
	read := solution(fr.getRead4())
	for _, tc := range []struct {
		ntoread int
		want    string
	}{
		{1, "a"},
		{4, "bc"},
		{1, ""},
	} {
		b := make([]byte, len(fr.contents))
		nread := read(b, tc.ntoread)
		b = b[:nread]
		require.Equal(t, tc.want, string(b))
	}
}

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	var prevBuf []byte
	return func(buf []byte, n int) int {
		ntotal := len(prevBuf)
		copy(buf, prevBuf)
		for ntotal < n {
			nread := read4(buf[ntotal:])
			ntotal += nread
			if nread < 4 {
				break
			}
		}
		// copy overflow into prevBuf
		if ntotal >= n {
			prevBuf = make([]byte, ntotal-n)
			copy(prevBuf, buf[n:ntotal])
		} else {
			prevBuf = prevBuf[:0]
		}
		return min(n, ntotal)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
