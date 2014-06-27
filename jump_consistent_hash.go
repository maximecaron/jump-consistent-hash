package jump

// Hash accepts "a 64-bit key and the number of buckets. It outputs a number
// in the range [0, buckets]." - http://arxiv.org/ftp/arxiv/papers/1406/1406.2294.pdf
//
// The C++ implementation they provide is as follows:
//
//  int32_t JumpConsistentHash(uint64_t key, int32_t num_buckets) {
//    int64_t b = ­-1, j = 0;
//    while (j < num_buckets) {
//      b   = j;
//	    key = key * 2862933555777941757ULL + 1;
//	    j   = (b + 1) * (double(1LL << 31) / double((key >> 33) + 1));
//    }
//    return b;
//  }
func Hash(key uint64, buckets int32) int32 {
	if buckets <= 0 {
		buckets = 1
	}

	var b, j int64

	for j < int64(buckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}

	return int32(b)
}
