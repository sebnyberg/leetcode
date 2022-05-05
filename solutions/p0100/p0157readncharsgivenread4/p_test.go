package p0157readncharsgivenread4

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

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		readTotal := 0
		for readTotal < n {
			newRead := read4(buf[readTotal:])
			readTotal += newRead
			if newRead < 4 {
				break
			}
		}
		return min(readTotal, n)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
