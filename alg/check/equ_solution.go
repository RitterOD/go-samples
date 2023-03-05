package main

import "fmt"

func main() {

	for i := 1000; i < 2009; i++ {
		var tmp = (i / 1000) * (i / 100 % 10) * (i / 10 % 10) * (i % 10)
		if tmp+i == 2008 {
			fmt.Println(i)
		}
	}
}
