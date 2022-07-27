# Compiler

This document contains snippets of code which will be optimized by the Go compiler.

## Go 1.11

### Clearing maps

```go
for k := range m {
  delete(m, k)
}
```

### Increasing the length of a slice

```go
s := make([]int, 100)
if cap(s) < 200 {
  // Need to extend s
  s = append(s, make([]int, 100)...)
}
```
