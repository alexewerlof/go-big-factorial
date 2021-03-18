# Intro

This is the good old factorial algorithm using Go's high performant [big number](https://golang.org/pkg/math/big/)
standard library.

## How does it work?

You run it like a CLI and give it a number larger than 2 like this:

```bash
go run . 100
```

It breaks every number to its prime factors. It does this for all numbers that are needed to be multiplied to calculate the factorial.
Then it computes all those powers and multiplications in go routines.
It's quite fast compared to traditional algorithms (and even my best implementation in Node.js).

1,000,000! takes less than a minute on a 10th gen i5 CPU with 8GB RAM on Windows 10.

# Debug

You can run this project in VS Code directly.
It gets its argument input from `./.vscode/launch.json` `arg` field. 

