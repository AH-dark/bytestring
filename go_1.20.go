//go:build go1.20 || go1.21 || go1.22
// +build go1.20 go1.21 go1.22

package bs

import "unsafe"

// S2B converts string to byte slice without a memory allocation.
// For more details, see https://github.com/golang/go/issues/53003#issuecomment-1140276077.
func S2B(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// B2S converts byte slice to string without a memory allocation.
// For more details, see https://github.com/golang/go/issues/53003#issuecomment-1140276077.
func B2S(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
