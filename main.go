package main

import (
	"fmt"

	hola1 "github.com/nicksnyder/hola/hola"
	hola2 "github.com/nicksnyder/hola/v2/hola"
)

func main() {
	fmt.Println(hola1.Hola())
	fmt.Println(hola2.Hola2())
}

