# prefixsuffix
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/ammario/prefixsuffix)

Exported [os/exec.prefixSuffixSaver](https://golang.org/src/os/exec/exec.go?s=19091:19652#L649).

```go
// Saver is an io.Writer which retains the first N bytes
// and the last N bytes written to it. The Bytes() methods reconstructs
// it with a pretty error message.
type Saver struct {
    N         int // max size of prefix or suffix
    // ...
}
```

## Basic Usage

```go
func something() {
    s := &prefixsuffix.Saver{N: 4}
    io.WriteString(s, "1234 --- 5678")
    fmt.Printf("%s\n", s.Bytes())   
    // 1234
    // ... omitting 5 bytes ...
    // 5678
}
```
