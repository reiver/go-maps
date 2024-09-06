# go-maps

Package **maps** provides tools for working with Go's `map[string]interface{}` type.

For example:
```go
var data map[string]interface{}

// ...

postalCode, found := maps.PathQuery(data, "user", "address", "postal-code")
```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-maps

[![GoDoc](https://godoc.org/github.com/reiver/go-maps?status.svg)](https://godoc.org/github.com/reiver/go-maps)

## Import

To import package **maps** use `import` code like the following:
```
import "github.com/reiver/go-maps"
```

## Installation

To install package **maps** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-maps
```

## Author

Package **maps** was written by [Charles Iliya Krempeaux](http://reiver.link)
