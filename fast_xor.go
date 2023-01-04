package byte_utils

// This file contains functions for xor-ing slices of bytes together.

import (
	"encoding/binary"
)

// Sets dst to the xor of bytes in slices a and b, one byte at a time. Panics
// if b and dst aren't at least as long as a. It is valid for dst to be the
// same as either a or b.
func SimpleXor(dst, a, b []byte) {
	for i := 0; i < len(a); i++ {
		dst[i] = a[i] ^ b[i]
	}
}

// Used to xor byte slices a and b, writing the result into dst. Based on
// github.com/golang/go/issues/31586#issuecomment-487436401. Panics if b and
// dst aren't at least as long as a (the length of the operation is determined
// by len(a).  It's valid for dst to be the same as either a or b. Benchmarks
// on my machine (tm) indicate that this is over 3x faster than SimpleXor alone
// for 1 MB slices, though this may change depending on architecture.
func FastXor(dst, a, b []byte) {
	n := binary.LittleEndian
	var x, y uint64
	for len(a) >= 32 {
		x = n.Uint64(a)
		y = n.Uint64(b)
		n.PutUint64(dst, x^y)
		x = n.Uint64(a[8:])
		y = n.Uint64(b[8:])
		n.PutUint64(dst[8:], x^y)
		x = n.Uint64(a[16:])
		y = n.Uint64(b[16:])
		n.PutUint64(dst[16:], x^y)
		x = n.Uint64(a[24:])
		y = n.Uint64(b[24:])
		n.PutUint64(dst[24:], x^y)
		a = a[32:]
		b = b[32:]
		dst = dst[32:]
	}
	SimpleXor(dst, a, b)
}
