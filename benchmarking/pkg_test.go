package main

import (
	"testing"
)

// function to Benchmark ReturnGeeks()
func BenchmarkGeeks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReturnGeeks()
	}
}

/*Command:

go test -bench=.
where -bench=. is flag need to run the default benchmark test.
You can manipulate different flags while testing.*/
