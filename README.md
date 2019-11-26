data_structure - Common Golang Data Structure
================================

[![Build Status](https://travis-ci.org/stretchr/testify.svg)](https://travis-ci.org/stretchr/testify) [![Go Report Card](https://goreportcard.com/badge/github.com/stretchr/testify)](https://goreportcard.com/report/github.com/stretchr/testify) [![GoDoc](https://godoc.org/github.com/stretchr/testify?status.svg)](https://godoc.org/github.com/stretchr/testify)

Go code (golang) set of packages that provide common data structure

------

Installation
============

To install data_structure, use `go get`:

    go get github.com/whuhyw/data_structure

This will then make the following packages available to you:

    github.com/whuhyw/data_structure/stack
    github.com/whuhyw/data_structure/queue

Import the `data_structure/stack` package into your code using this template:

```go
package main

import (
  "testing"

  "github.com/whuhyw/data_structure/stack"
)

func main() {
    s := stack.NewStack()
    
    s.Push("hello")
    if !s.Empty(){
        s.Pop()
    }
}
```

------

Staying up to date
==================

To update Testify to the latest version, use `go get -u github.com/whuhyw/data_structure`.

------

Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.

------

License
=======

This project is licensed under the terms of the MIT license.