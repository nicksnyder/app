package main

import (
	"fmt"

	h1 "github.com/nicksnyder/hello/hello"
	h2 "github.com/nicksnyder/hello/v2/hello"
)

func main() {
	fmt.Println(h1.Hello())
	fmt.Println(h2.Hello2())
}
