# gomovave

This package implements a moving average function implementation in Go with the computed `Value()` method.

## Requirements

* go 1.18+

## Installation

```sh
go get github.com/Devoter/gomovave
```

## Usage

### Creating an instance

```go
import "github.com/Devoter/gomovave"

const maxSize = 10

ma := gomovave.NewMovingAverage[float64](maxSize)
```

### Appending values

```go
ma.Push(6)
ma.Push(3)
ma.Push(3)
```

### Reading a value

`Value()` method is a computed and it isn't calculated twice without changing the queue.

```go
ma.Value() // 4
```

### Getting the current queue length

```go
ma.Len() // 3
```

### Getting the maximum queue length

```go
ma.MaxLen() // 10
```

### Getting a front value

```go
ma.Front() // 6
```

### Getting a raw queue

**ATTENTION**: `Queue()` method returns a raw queue pointer due to the performance.

```go
q := ma.Queue() // [6, 3, 3]
```

### Cleaning the queue

```go
ma.Clear()
ma.Len() // 0
```

## License

[LICENSE](./LICENSE)
