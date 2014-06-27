package jump

import (
	"testing"
)

func TestHashInBucketRange(t *testing.T) {
	ch := Hash(0, 10)
	if ch < 0 || ch > 10 {
		t.Error("expected bucket to be in range [0, 10]; got", ch)
	}
}

func TestNegativeBucketsEqualsOneBucket(t *testing.T) {
	ch := Hash(0, -10)
	if ch < 0 || ch > 1 {
		t.Error("expected bucket to be in range [0, 1]; got", ch)
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hash(uint64(i), int32(i))
	}
}
