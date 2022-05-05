package p0535encodeanddecodetinyurl

import (
	"encoding/base64"
	"encoding/binary"
	"hash"
	"hash/fnv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Codec(t *testing.T) {
	c := Constructor()
	url := "https://leetcode.com/problems/design-tinyurl"
	encoded := c.encode(url)
	decoded := c.decode(encoded)
	require.Equal(t, url, decoded)
}

type Codec struct {
	h    hash.Hash32
	urls map[uint32]string
}

func Constructor() Codec {
	return Codec{
		h:    fnv.New32a(),
		urls: make(map[uint32]string),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	b := []byte(longUrl)
	key := this.h.Sum32()
	keyBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(keyBytes, key)
	defer this.h.Reset()
	this.urls[key] = base64.RawURLEncoding.EncodeToString(b)
	return "http://tinyurl.com/" + base64.RawURLEncoding.EncodeToString(keyBytes)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	parts := strings.Split(shortUrl, "/")
	kBytes := parts[len(parts)-1]
	decodedKeyBytes, _ := base64.RawURLEncoding.DecodeString(kBytes)
	key := binary.BigEndian.Uint32(decodedKeyBytes)
	encodedURL := this.urls[key]
	b, _ := base64.RawURLEncoding.DecodeString(encodedURL)
	return string(b)
}
