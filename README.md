# experiment-inlined-pointer
This repository is for an experiment of inlined pointers.

## example
```Shell
% make benchmark-comparison
##### Benchmarking... #####
goos: darwin
goarch: arm64
pkg: github.com/YukiBobier/experiment-inlined-pointer
BenchmarkPrintFoo-10            15016969                79.77 ns/op            0 B/op          0 allocs/op
BenchmarkPrintFooPointer-10     15988239                74.67 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/YukiBobier/experiment-inlined-pointer        3.814s
##### Benchmarking without inlining... #####
goos: darwin
goarch: arm64
pkg: github.com/YukiBobier/experiment-inlined-pointer
BenchmarkPrintFoo-10            12429102                97.24 ns/op            0 B/op          0 allocs/op
BenchmarkPrintFooPointer-10      7213098               147.8 ns/op          1024 B/op          1 allocs/op
PASS
ok      github.com/YukiBobier/experiment-inlined-pointer        3.776s
```
