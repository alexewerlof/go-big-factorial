# Intro

This is the good old factorial algorithm using Go's high performant [big number](https://golang.org/pkg/math/big/)
standard library.

# Compile

Run `go build` in the root.

# Execute

When you build it, you get an executable file.
Run it passing a number. Start with something small because numbers bigger than 10000 can take a while.

```bash
go-big-factorial 300
```

# Debug

You can run this project in VS Code directly.
It gets its argument input from `./.vscode/launch.json` `arg` field. 

