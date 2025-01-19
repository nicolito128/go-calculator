# Go Calculator
A simple calculator module written in Go.

It accept the following operations: `+` (addition), `-` (subtraction), `*` (product), `/` (division) and `^` (power). Also, you can use parenthesis association `()`.

## Getting started

Get the module with:

    go get github.com/nicolito128/go-calculator

Import the module:

```go
    import "github.com/nicolito128/go-calculator"
```

Now you can use the `Resolve function`. For example:

```go
    package main

    import (
        "fmt"

        "github.com/nicolito128/go-calculator"
    )

    func main() {
        result, err := calculator.Resolve("100 + ((2 / 4) * 2 * 3)")
        if err != nil {
            panic("Oh no! An error accured.")
        }
        
        fmt.Println(result)
    }
```
