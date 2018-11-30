# Merkle Tree Benchmark

This program provide a benchmark tool to returns metrics about a merkle tree.
The metrics are:

- How many seconds to build a merkle tree within N transactions
- How many seconds to verify the merkle tree within N transactions
- How many seconds to check if a transaction belongs to the merkle tree within N transactions

## How to use

```go
go run main.go -n {NUMBER_OF_TRANSACTIONS}
```

## Benchmarks

| Transactions number |Tree building|Tree verifying|Transaction existing|
|:-|:-|:-|:-|
| 1000 | 0.001405 seconds | 0.001922 seconds | 0.000040 seconds |
| 10000 | 0.013499 seconds | 0.020007 seconds | 0.000330 seconds |
| 100000 | 0.153391 seconds | 0.175260 seconds | 0.002227 seconds |
| 1000000 | 1.425276 seconds | 1.336388 seconds | 0.032563 seconds |
