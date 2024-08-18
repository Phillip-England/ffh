# ffh
read and write `.go` files

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
```

## Examples

create a `.go` func as a string and write it out to a file
```go
func main() {
    path := "./some.go"

    // make a file
    err := ffh.Touch(path)
    if err != nil {
        panic(err)
    }

    // create a go func as a string
    goFunc, err := ffh.GoFunc("main", "", "", ffh.LinesToStr(
        "fmt.Println(\"Hello, World!\")",
    ))

    // write it to a file
    err := ffh.OverwriteFile(path, goFunc)
    if err != nil {
        panic(err)
    }

    fmt.Println(goFunc)

}
```