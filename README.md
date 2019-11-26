container - Common Golang Data Structure
================================

[![Build Status](https://travis-ci.org/whuhyw/container.svg)](https://travis-ci.org/whuhyw/container) [![Go Report Card](https://goreportcard.com/badge/github.com/whuhyw/container)](https://goreportcard.com/report/github.com/whuhyw/container) [![GoDoc](https://godoc.org/github.com/whuhyw/container?status.svg)](https://godoc.org/github.com/whuhyw/container)

Go code (golang) set of packages that provide common data structure

------

Installation
============

To install container, use `go get`:

    go get github.com/whuhyw/container

This will then make the following packages available to you:

    github.com/whuhyw/container/stack
    github.com/whuhyw/container/queue

Import the `container/stack` package into your code using this template:

```go
package main

import (
  "github.com/whuhyw/container/stack"
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

To update container to the latest version, use `go get -u github.com/whuhyw/data_structure`.

------

Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.

------

License
=======

This project is licensed under the terms of the MIT license.