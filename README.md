# Go Versions

[![BuildStatus](https://github.com/cappyzawa/go-versions/workflows/CI/badge.svg)](https://github.com/cappyzawa/go-versions/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/cappyzawa/go-versions)](https://goreportcard.com/report/github.com/cappyzawa/go-versions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/cappyzawa/go-versions)](https://pkg.go.dev/github.com/cappyzawa/go-versions)
[![codecov](https://codecov.io/gh/cappyzawa/go-versions/branch/master/graph/badge.svg)](https://codecov.io/gh/cappyzawa/go-versions)

This tool gets go versions from [Downloads \- The Go Programming Language](https://golang.org/doc/).

## How to use

```bash
# no fileter
go-versions

# fileter: os=linux
go-versions -os linux

# fileter: arch=amd64
go-versions -arch amd64

# fileter: os=linux arch=amd64
go-versions -os linux -arch amd64
```
