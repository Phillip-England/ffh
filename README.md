# ffh
ffh is a minimal package for extracting details from `.go` files

## Installation
```bash
go get github.com/Phillip-England/ffh
```

## Import
```go
package main

import (
    "fmt"
    "github.com/Phillip-England/ffh"
)

func main() {
    str, err := ffh.ReadFile("./main.go")
    if err != nil {
        panic(err)
    }
}
```

## Usage

get the package line out of a `.go` file
```go
func main() {
    str, err := ffh.ReadFile("./main.go")
    if err != nil {
        panic(err)
    }
    packageLine, err := ffh.ExtractPackageLine(str)
    if err != nil {
        panic(err)
    }
    fmt.Println(packageLine)
}
```

get the import block out of a `.go` file
```go
func main() {
    str, err := ffh.ReadFile("./main.go")
    if err != nil {
        panic(err)
    }
    importBlock, err := ffh.ExtractImportBlock(str)
    if err != nil {
        panic(err)
    }
    fmt.Println(importBlock)
}
```

get the funcs out of a `.go` file
```go
func main() {
    str, err := ffh.ReadFile("./main.go")
    if err != nil {
        panic(err)
    }
    goFuncs, err := ffh.ExtractFuncBlocks(str)
    if err != nil {
        panic(err)
    }
    for _, fn := range goFuncs {
        fmt.Println(fn)
    }
}
```

get the type definitions out of a `.go` file
```go
type TestType struct {
    HasPersonality bool // 🤩
}


func main() {
    str, err := ffh.ReadFile("./main.go")
    if err != nil {
        panic(err)
    }
    goTypes, err := ffh.ExtractTypeBlocks(str)
    if err != nil {
        panic(err)
    }
    for _, goType := range goTypes {
        fmt.Println(goType)
    }
}
```