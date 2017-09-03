package main

import "fmt"

func main() {

	a , b :=100,200
	var c * int

	c = &a

	fmt.Printf("a  , b ,c  %T, %T, %T\n", a , &b ,c)
	fmt.Printf("a  , b ,c  %d, %d, %d\n", &a , &b ,c)
	fmt.Printf("a  , b ,c  %d, %d, %d\n", a , b ,*c)
}