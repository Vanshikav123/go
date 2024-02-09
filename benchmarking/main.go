package main

//each output against a standard , assessing the code's overall performance level
//func Benchmark(b* testing.B)
/*In automation testing, most of the frameworks support only one among functionality testing and
 benchmark testing. But the Golang testing package provides many functionalities for a different type of testing including benchmark testing.

B is a type(struct) passed to Benchmark functions to manage benchmark timing and to specify the number of iterations to run.
Basically benchmark test suite of the testing package gives the benchmark reports likes time consumed, number of iteration/request(i.e. execution of function) of the tested function.

Syntax:

func BenchmarkXxx(*testing.B)
All of the benchmark function is executed by go test command. BenchmarkResult contains the results of a benchmark run.
type BenchmarkResult struct {
    N         int           // The number of iterations.
    T         time.Duration // The total time taken.
    Bytes     int64         // Bytes processed in one iteration.
    MemAllocs uint64        // The total number of memory allocations; added in Go 1.1
    MemBytes  uint64        // The total number of bytes allocated; added in Go 1.1

    // Extra records additional metrics reported by ReportMetric.
    Extra map[string]float64 // Go 1.13
}*/

// function which return "geeks"
func ReturnGeeks() string {
	return "geeks"
}

// main function of package
func main() {
	ReturnGeeks()
}
