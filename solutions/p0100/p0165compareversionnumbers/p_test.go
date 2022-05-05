package p0165compareversionnumbers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_compareVersion(t *testing.T) {
	for _, tc := range []struct {
		version1 string
		version2 string
		want     int
	}{
		{"1.01", "1.001", 0},
		{"1.0", "1.0.0", 0},
		{"0.1", "1.1", -1},
		{"3.0.4.10", "3.0.4.2", 1},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.version1, tc.version2), func(t *testing.T) {
			require.Equal(t, tc.want, compareVersion(tc.version1, tc.version2))
		})
	}
}

func compareVersion(version1 string, version2 string) int {
	v1 := parseVersion(version1)
	v2 := parseVersion(version2)
	return v1.compare(v2)
}

func MustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return n
}

func parseVersion(s string) version {
	parts := strings.Split(s, ".")
	var v version
	v.major = MustParseInt(parts[0])
	if len(parts) > 1 {
		v.minor = MustParseInt(parts[1])
	}
	if len(parts) > 2 {
		v.patch = MustParseInt(parts[2])
	}
	if len(parts) > 3 {
		v.patchpatch = MustParseInt(parts[3])
	}
	return v
}

type version struct {
	major      int
	minor      int
	patch      int
	patchpatch int
}

func (v version) compare(other version) int {
	switch {
	case v.major < other.major,
		v.major == other.major && v.minor < other.minor,
		v.major == other.major && v.minor == other.minor && v.patch < other.patch,
		v.major == other.major && v.minor == other.minor && v.patch == other.patch && v.patchpatch < other.patchpatch:
		return -1
	case v.major > other.major,
		v.major == other.major && v.minor > other.minor,
		v.major == other.major && v.minor == other.minor && v.patch > other.patch,
		v.major == other.major && v.minor == other.minor && v.patch == other.patch && v.patchpatch > other.patchpatch:
		return 1
	default:
		return 0
	}
}
