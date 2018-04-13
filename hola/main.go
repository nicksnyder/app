package main

import (
	"fmt"

	hello1 "github.com/nicksnyder/hola/hola"
	hello2 "github.com/nicksnyder/hola/v2/hola"
)

func main() {
	fmt.Println(hola1.Hola())
	fmt.Println(hola2.Hola())
}

