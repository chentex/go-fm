# go-fm

A simple file manager library.

[![Build Status](https://travis-ci.org/chentex/go-fm.svg)](https://travis-ci.org/chentex/go-fm)
[![codecov](https://codecov.io/gh/chentex/go-fm/branch/master/graph/badge.svg)](https://codecov.io/gh/chentex/go-fm/branch/master)
[![GoDoc](https://godoc.org/github.com/chentex/go-fm?status.svg)](https://godoc.org/github.com/chentex/go-fm)
[![Go Report Card](https://goreportcard.com/badge/github.com/chentex/go-fm)](https://goreportcard.com/report/github.com/chentex/go-fm)

Package errors provides simple error handling primitives and behavioral errors.

`go get github.com/chentex/go-fm`

[Read the package documentation for more information](https://godoc.org/github.com/chentex/go-fm).

## Motivation

Every time you need to manage some files you tend to use the same code over and over, and you always have to test no matter how many times you have written the code.

This library is a wrapper for reading, writing and checking if files exists. It's already tested so you can use it without worrying to write the test for this module, and just focus on writing your code.

## Usage

### Import

`import "github.com/chentex/go-fm" fm`

### New File Manager

```
fileManager := fm.NewFileManager()
```

### Reading a file

```
content, err := fm.OpenFile("yourfile.txt")
if err != nil {
    fmt.Printf(err)
}
```

### Writing a file

```
bytes := byte[]("sample text to insert in file")
err := fm.WriteFile("yourfile.txt", bytes, 0644)
if err != nil {
    fmt.Printf(err)
}
```

### Checking for a file

```
exists, err := fm.ExistsFile("yourfile.txt")
if err != nil {
    fmt.Printf(err)
}

fmt.Printf(exists)
```

## Contributing

I welcome pull requests, bug fixes and issue reports.

Maintainer: Vicente Zepeda (chente.z.m@gmail.com)