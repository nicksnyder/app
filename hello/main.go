package main

import (
	"fmt"

	hello1 "github.com/nicksnyder/hello/hello"
	hello2 "github.com/nicksnyder/hello/v2/hello"

)

func main() {
	fmt.Println(hello1.Hello())
	fmt.Println(hello2.Hello2())
}

