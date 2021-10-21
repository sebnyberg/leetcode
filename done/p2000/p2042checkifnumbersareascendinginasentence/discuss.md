# Check if numbers are ascending in a sentence

## Clean, idiomatic

```go
func areNumbersAscending(s string) bool {
	parts := strings.Split(s, " ")
	prev := -1
	for _, part := range parts {
		if n, err := strconv.Atoi(part); err == nil {
			if n <= prev {
				return false
			}
			prev = n
		}
	}
	return true
}
```

## Optimized

```go
func areNumbersAscending(s string) bool {
	prev := -1
	var cur int
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			cur *= 10
			cur += int(ch - '0')
			continue
		}
		if cur != 0 {
			if cur <= prev {
				return false
			}
			prev = cur
			cur = 0
		}
	}
	return cur == 0 || prev < cur
}
```
