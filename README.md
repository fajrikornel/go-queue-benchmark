
# Go Queue Benchmark

This project benchmarks two type of queue implementations in Go: Array queue and List queue.

To be noted, the array implementation uses a fixed-size Go slice internally. This is so that the queue capacity can be configurable at runtime.

## Run instructions

Test:
```
make test
```

Run benchmark:
```
make bench
```

## Results

On my laptop (M1 chip), the results are:
```
goos: darwin
goarch: arm64
pkg: github.com/fajrikornel/go-queue-benchmark/cmd
BenchmarkArrayQueue_Enqueue-8           1000000000               0.004126 ns/op
BenchmarkListQueue_Enqueue-8            1000000000               0.03068 ns/op
BenchmarkArrayQueue_Dequeue-8           1000000000               0.002648 ns/op
BenchmarkListQueue_Dequeue-8            1000000000               0.002811 ns/op
PASS
ok      github.com/fajrikornel/go-queue-benchmark/cmd   0.629s
```
Most attempts have similar results.

It seems that:
- Dequeue performance for both array & list queue are most similar, with an average difference of ~0.0002 ns (array is faster).
- For enqueue performance, array queue is faster than list queue by an average difference of ~0.02-0.03 ns.
