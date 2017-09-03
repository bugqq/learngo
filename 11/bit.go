package main

import "fmt"

func main() {

	var aa uint = 10011001
	var bb uint = 01100110
	var cc uint
	cc = aa & bb
   fmt.Println("c  = %d", cc)

}