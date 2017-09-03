package main

import "fmt"

func main() {
	var x interface{}
 	test(x)

	a:=100
	x = a
	test(x)
	
	b:="test"
	x=b
	test(x)
	
}

func test (b  interface{}) ( ) {
	switch b.(type) {
		case nil:
			fmt.Println("is nil")
		case int:
			fmt.Println("is int")
		case func(int):
			fmt.Println("is func int")
		case bool,string:
			fmt.Println("is bool or string")
		default:
			fmt.Println("default")
	}
}
