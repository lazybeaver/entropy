// Tests for Shannon Entropy

package entropy

import (
	"math"
	"testing"
)

func approxEqual(lhs, rhs float64) bool {
	return math.Abs(lhs-rhs) < 1e-14
}

func getMaxEntropyString() string {
	buf := make([]byte, 256)
	for i := 0; i < 256; i++ {
		buf[i] = byte(i)
	}
	return string(buf)
}

func verifyEntropy(t *testing.T, str string, expected float64) {
	actual, err := Shannon(str)
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err)
	}
	if !approxEqual(expected, actual) {
		t.Errorf("%s: expected=%g actual=%g", str, expected, actual)
	}
}

func TestMinShannonEntropy(t *testing.T) {
	verifyEntropy(t, "", 0.0)
	verifyEntropy(t, "a", 0.0)
	verifyEntropy(t, "aaaaa", 0.0)
}

func TestMaxShannonEntropy(t *testing.T) {
	verifyEntropy(t, getMaxEntropyString(), 8.0)
}

func TestOtherShannonEntropy(t *testing.T) {
	verifyEntropy(t, "ab", 1.0)
	verifyEntropy(t, "aabbccdd", 2.0)
	verifyEntropy(t, "abcdefghijklmnopqrstuvwxyz012345", 5.0)
	verifyEntropy(t, "abbcccdddd", 1.8464393446710154)
	verifyEntropy(t, "Hello World", 2.8453509366224368)
}

func BenchmarkShannon(b *testing.B) {
	s := getMaxEntropyString()
	for n := 0; n < b.N; n++ {
		Shannon(s)
	}
}
