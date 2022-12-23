package main

import (
	"fmt"
	"go-samples/hello"
)

func main() {
	message := hello.Hello("Gladys")
	fmt.Println(message)
}
