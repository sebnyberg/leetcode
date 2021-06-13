package unsafex

// NOTE!
// This code is taken from https://github.com/bcmills/unsafeslice/blob/master/unsafeslice.go
//
// The reason for putting it here is so that I can easily copy/paste
// the solution for use in leetcode challenges
//

import (
	"reflect"
	"unsafe"
)

// OfString returns a slice that refers to the data backing the string s.
//
// The caller must ensure that the contents of the slice are never mutated.
//
// Programs that use OfString should be tested under the race detector to flag
// erroneous mutations.
//
// Programs that have been adequately tested and shown to be safe may be
// recompiled with the "unsafe" tag to significantly reduce the overhead of this
// function, at the cost of reduced safety checks. Programs built under the race
// detector always have safety checks enabled, even when the "unsafe" tag is
// set.
func OfString(s string) []byte {
	p := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)

	var b []byte
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	hdr.Data = uintptr(p)
	hdr.Cap = len(s)
	hdr.Len = len(s)

	// maybeDetectMutations(b)
	return b
}

// AsString returns a string that refers to the data backing the slice s.
//
// The caller must ensure that the contents of the slice are never again
// mutated, and that its memory either is managed by the Go garbage collector or
// remains valid for the remainder of this process's lifetime.
//
// Programs that use AsString should be tested under the race detector to flag
// erroneous mutations.
//
// Programs that have been adequately tested and shown to be safe may be
// recompiled with the "unsafe" tag to significantly reduce the overhead of this
// function, at the cost of reduced safety checks. Programs built under the race
// detector always have safety checks enabled, even when the "unsafe" tag is
// set.
func AsString(b []byte) string {
	p := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)

	var s string
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(p)
	hdr.Len = len(b)

	// maybeDetectMutations(b)
	return s
}
