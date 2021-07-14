# Intro

This is a very fast algorithm ([prime factorization](https://janmr.com/blog/2010/10/prime-factors-of-factorial-numbers/)) and CLI using Go's high performance [big number](https://golang.org/pkg/math/big/)
standard library together with go routines.

## How does it work?

You run it like a CLI and give it a number larger than 2 like this:

```bash
go run . 100
```

It breaks every number to its prime factors. It does this for all numbers that are needed to be multiplied to calculate the factorial.
Then it computes all those powers and multiplications in go routines.
It's quite fast compared to traditional algorithms (and even my best implementation in Node.js).

1,000,000! takes less than a minute on a 10th gen i5 CPU with 8GB RAM on Windows 10.

## Test

```
go test .
```

## Benchmark

Breaking down the text above, we pass the -bench flag to go test supplying a regular expression matching everything. You must pass a valid regex to -bench, just passing -bench is a syntax error. You can use this property to run a subset of benchmarks.

[ref](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

```
go test -bench=.
```

# Debug

You can run this project in VS Code directly.
It gets its argument input from `./.vscode/launch.json` `arg` field. 
