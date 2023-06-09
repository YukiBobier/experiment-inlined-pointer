# experiment-inlined-pointer
This repository is for an experiment of inlined pointers.

## example
```Shell
% make benchmark-comparison
##### Benchmarking... #####
goos: darwin
goarch: arm64
pkg: github.com/YukiBobier/experiment-inlined-pointer
BenchmarkPrintFoo-10               46504             24772 ns/op               0 B/op          0 allocs/op
BenchmarkPrintFooPointer-10        48526             24812 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/YukiBobier/experiment-inlined-pointer        3.176s
##### Benchmarking without inlining... #####
goos: darwin
goarch: arm64
pkg: github.com/YukiBobier/experiment-inlined-pointer
BenchmarkPrintFoo-10               44491             26120 ns/op               0 B/op          0 allocs/op
BenchmarkPrintFooPointer-10        45529             26380 ns/op           65536 B/op          1 allocs/op
PASS
ok      github.com/YukiBobier/experiment-inlined-pointer        3.129s
```
