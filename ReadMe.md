# Testing
```
    go test ./...
    go test -v ./...
    go test -v ./... -run TestName
```

# Benchmarking
```
    go test -bench=.
    go test -bench=. -benchmem
    go test -bench=. -benchmem -benchtime=10s
```