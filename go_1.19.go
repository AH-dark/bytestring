//go:build !go1.20 && !go1.21 && !go1.22
// +build !go1.20,!go1.21,!go1.22

package bs

import "unsafe"

// S2B converts string to byte slice without a memory allocation.
func S2B(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// B2S converts byte slice to string without a memory allocation.
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
