# Merkle Tree Benchmark

This program provide a benchmark tool to returns metrics about a merkle tree.

## How to use

```go
go run main.go -n {NUMBER_OF_TRANSACTIONS}
```

## Benchmarks

| Transactions number |Tree building|Deep tree verifying (validates the hashes at each level of the tree and compare with the root of the tree matches) |Verify transaction integrity (calculate on the critical path and compare with the merkle tree root)|
|:-|:-|:-|:-|
| 1000 | 0.001405 seconds | 0.001922 seconds | 0.000040 seconds |
| 10000 | 0.013499 seconds | 0.020007 seconds | 0.000330 seconds |
| 100000 | 0.153391 seconds | 0.175260 seconds | 0.002227 seconds |
| 1000000 | 1.425276 seconds | 1.336388 seconds | 0.032563 seconds |
