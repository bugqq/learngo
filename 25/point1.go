package main

import "fmt"

func main() {

	var a *int
	var b *float32
	  fmt.Printf("%x\n", a)
	fmt.Printf("%x\n", b)
	
	c:=1
	a=&c
	var d float32 = 10.0
	b = &d
	fmt.Printf("%x\n", a)
	fmt.Printf("%x\n", b)
	
}