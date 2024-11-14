# GOHELPER

**Maintainer:** jiharal

![Go CI](https://github.com/tlabdotcom/gohelper/actions/workflows/go.yml/badge.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/tlabdotcom/gohelper)](https://goreportcard.com/report/github.com/tlabdotcom/gohelper)

## Overview

The `GOHELPER` package provides utility functions for working with pointers and safely retrieving values from maps in Go. These helper functions can make your code cleaner and more concise.

## Functions

### Pointer Functions

These functions return pointers to values of their respective types. They are useful for creating pointer references for various data types.

#### Available Functions

- **PointerUUID**: Returns a pointer to a `uuid.UUID`.
- **PointerString**: Returns a pointer to a `string`.
- **PointerBool**: Returns a pointer to a `bool`.
- **PointerInt**: Returns a pointer to an `int`.
- **PointerInt8**: Returns a pointer to an `int8`.
- **PointerInt16**: Returns a pointer to an `int16`.
- **PointerInt32**: Returns a pointer to an `int32`.
- **PointerInt64**: Returns a pointer to an `int64`.
- **PointerUint**: Returns a pointer to a `uint`.
- **PointerUint8**: Returns a pointer to a `uint8`.
- **PointerUint16**: Returns a pointer to a `uint16`.
- **PointerUint32**: Returns a pointer to a `uint32`.
- **PointerUint64**: Returns a pointer to a `uint64`.
- **PointerFloat32**: Returns a pointer to a `float32`.
- **PointerFloat64**: Returns a pointer to a `float64`.
- **PointerComplex64**: Returns a pointer to a `complex64`.
- **PointerComplex128**: Returns a pointer to a `complex128`.
- **PointerByte**: Returns a pointer to a `byte`.
- **PointerRune**: Returns a pointer to a `rune`.

#### Usage Example

```shell
go get github.com/tlabdotcom/gohelper
```

```go
package main

import (
    "fmt"
    "github.com/google/uuid"
    "github.com/tlabdotcom/gohelper"
)

func main() {
    str := "Hello, World!"
    strPtr := gohelper.PointerString(str)
    fmt.Println(*strPtr) // Output: Hello, World!

    id := uuid.New()
    idPtr := gohelper.PointerUUID(id)
    fmt.Println(*idPtr) // Output: (some UUID value)
}
```
