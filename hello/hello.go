package main

import "fmt"
import "sample/hello/randomstr"

func main() {
	randomMessage := randomstr.RandomStr(100)
	fmt.Printf("our random message is %s", randomMessage)
}
