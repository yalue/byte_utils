package byte_utils

import (
	"bytes"
	"math/rand"
	"testing"
)

func randomBytes(size, seed int) []byte {
	rng := rand.New(rand.NewSource(int64(seed)))
	toReturn := make([]byte, size)
	for i := range toReturn {
		toReturn[i] = byte(rng.Int())
	}
	return toReturn
}

func TestXor(t *testing.T) {
	x := randomBytes(1337, 7)
	y := randomBytes(len(x), 8)
	expectedResult := make([]byte, len(x))
	for i := range x {
		expectedResult[i] = x[i] ^ y[i]
	}
	dst := make([]byte, len(x))
	SimpleXor(dst, x, y)
	if !bytes.Equal(dst, expectedResult) {
		t.Logf("SimpleXor produced incorrect results.\n")
		t.FailNow()
	}
	FastXor(dst, dst, y)
	if !bytes.Equal(dst, x) {
		t.Logf("FastXor failed to restore original input.\n")
		t.FailNow()
	}
}

func BenchmarkFastXor(b *testing.B) {
	size := 1024 * 1024
	x := randomBytes(size, 1337)
	y := randomBytes(size, 1338)
	dst := make([]byte, size)
	for n := 0; n < b.N; n++ {
		FastXor(dst, x, y)
	}
}

func BenchmarkSimpleXor(b *testing.B) {
	size := 1024 * 1024
	x := randomBytes(size, 1337)
	y := randomBytes(size, 1338)
	dst := make([]byte, size)
	for n := 0; n < b.N; n++ {
		SimpleXor(dst, x, y)
	}
}
