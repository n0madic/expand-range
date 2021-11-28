# expand-range

[![Go Reference](https://pkg.go.dev/badge/github.com/n0madic/expand-range.svg)](https://pkg.go.dev/github.com/n0madic/expand-range)

Parsing of numeric ranges from string.

Convert `"1,3-5"` into `[1,3,4,5]`.

## Installation

```shell
go get -u github.com/n0madic/expand-range
```

## Usage

Import:

```go
import expandrange "github.com/n0madic/expand-range"
```

Parse string:

```go
rng, err := expandrange.Parse("1,3-5")
if err != nil {
    panic(err)
}
```

Range checking:

```go
if rng.InRange(4) {
    ...
}
```

Sorting result:

```go
rng, _ := expandrange.Parse("3,2,1")
rng.Sort()
// [1,2,3]
```
