# go-mapstringinterface

Package **mapstringinterface** provides tools for working with Go's `map[string]interface{}` type.

For example:
```go
var data map[string]interface{}

// ...

postalCode, found := mapstringinterface.PathQuery(data, "user", "address", "postal-code")
```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-mapstringinterface

[![GoDoc](https://godoc.org/github.com/reiver/go-mapstringinterface?status.svg)](https://godoc.org/github.com/reiver/go-mapstringinterface)
