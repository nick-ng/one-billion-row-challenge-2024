# one-billion-row-challenge-2024
[1 Billion Row Challenge](https://github.com/gunnarmorling/1brc) but in Go because I don't know Java. So I guess not really _the_ 1 Billion Row Challenge. Just _a_ 1 billion row challenge?

1. `python make-data.py 1_000_000_000`
2. `go build`
3. run the binary. e.g. on Windows: `./one-billion-row-challenge.exe`

## Different Builds

- `main.go.01.stop` - implementation with my "hashmap"
- `main.go.02.stop` - implementation with Go's map
