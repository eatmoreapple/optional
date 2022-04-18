# Option

Optional is a library that provides option types for Go types.

## Installation

Golang version 1.11 + required

```shell
go get -u github.com/eatmoreapple/optional
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/eatmoreapple/optional"
)

func main() {
	a := optional.Some[int](1)
	fmt.Println(a)
	fmt.Println(a.IsNone())
	fmt.Println(a.IsSome())
	fmt.Println(a.Unwrap())
	fmt.Println(a.UnwrapOr(2))
	b := a.Replace(2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a.Map(func(i int) int { return i + 1 }))
	fmt.Println(a.Filter(func(i int) bool { return i == 2 }))
	fmt.Println(optional.None[string]().IsNone())
}
```
