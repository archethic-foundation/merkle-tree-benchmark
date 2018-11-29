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