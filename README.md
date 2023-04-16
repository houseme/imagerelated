# Image-Related

[![Go Reference](https://pkg.go.dev/badge/github.com/houseme/imagerelated.svg)](https://pkg.go.dev/github.com/houseme/imagerelated)
[![Action-CI](https://github.com/houseme/imagerelated/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/houseme/imagerelated/actions/workflows/go.yml)
![GitHub](https://img.shields.io/github/license/houseme/imagerelated?style=flat-square)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/houseme/imagerelated/main?style=flat-square)

Picture type information and get picture-related suffixes

## Install

```bash
go get -u -v github.com/houseme/imagerelated@main
```

## Usage

```go
package main

import (
    "fmt"
    
    "github.com/houseme/imagerelated"
)

func main() {
    // Get the suffix of the picture
    fmt.Println(imagerelated.ImageExt("testdata/golang.png"))
    // Get the type of the picture
    fmt.Println(imagerelated.ImageType("testdata/golang.png"))
}
```

## License
FeiE is primarily distributed under the terms of both the [Apache License (Version 2.0)](LICENSE)
