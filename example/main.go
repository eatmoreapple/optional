package main

import (
	"fmt"
	"github.com/eatmoreapple/optional"
)

func main() {
	hello := optional.Some[string]("Hello")
	world := optional.Some[string](" World")
	result := world.Map(func(s string) string {
		return hello.Unwrap() + s
	})
	fmt.Println(result.Unwrap())
}
