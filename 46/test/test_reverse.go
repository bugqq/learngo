package main

import (
	"fmt"
	"test2"
)

func main() {
	input := "Hello world"
	fmt.Println(input)
	ret := test.Reverse(input)
	fmt.Println(ret)
}
