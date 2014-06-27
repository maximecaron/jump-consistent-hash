# Jump Consistent Hash

[![GoDoc](https://godoc.org/github.com/anachronistic/jump-consistent-hash?status.png)](https://godoc.org/github.com/anachronistic/jump-consistent-hash)

A Go implementation of the jump consistent hash from Lamping and Veach.

Paper: [http://arxiv.org/ftp/arxiv/papers/1406/1406.2294.pdf](http://arxiv.org/ftp/arxiv/papers/1406/1406.2294.pdf)

## Usage

To use, simply import `"github.com/anachronistic/jump-consistent-hash"`  and call `Hash`, passing a `uint64` key and an `int32` number of buckets; the function will return an `int32`, as shown below.

```go
package main

import (
    jump "github.com/anachronistic/jump-consistent-hash"
)

func main() {
    jump.Hash(1, 1)      // 0
    jump.Hash(256, 1024) // 520
    jump.Hash(42, 57)    // 43
}
```

Note that if you pass a number of buckets <= 0 a default value of 1 will be used instead.

```go
jump.Hash(0xDEAD10CC, -666) // 0
jump.Hash(0xDEAD10CC, 1)    // 0
jump.Hash(0xDEAD10CC, 666)  // 361
```