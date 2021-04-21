package p0271encodeanddecodestrings

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"testing"
)

func TestCodec(t *testing.T) {
	c := Codec{}
	res := c.Encode([]string{"a", "b", "c"})
	fmt.Println(res)
	strs := c.Decode(res)
	fmt.Printf("%+v\n", strs)
}

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	res := make([]byte, 0)
	for _, s := range strs {
		n := len(s)
		buf := make([]byte, 8)
		nw := binary.PutUvarint(buf, uint64(n))
		res = append(res, buf[:nw]...)
		res = append(res, []byte(s)...)
	}
	return string(res)
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	res := make([]string, 0)
	b := bytes.NewBufferString(strs)
	for {
		strWidth, err := binary.ReadUvarint(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		buf := make([]byte, strWidth)
		_, err = b.Read(buf)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, string(buf))
	}
	return res
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
