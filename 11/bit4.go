package main

import "fmt"

func main() {

	var aa int = 10011001
	var bb int = 01100110
	var cc int
	cc = aa & bb
   fmt.Printf("c  等于  %d\n", cc)

   	cc = aa | bb
   fmt.Printf("c  等于  %d\n", cc)

	cc = aa ^ bb
   fmt.Printf("c  等于  %d\n", cc)

   cc = aa <<1
   fmt.Printf("c  等于  %d\n", cc)

   cc = aa >>1
   fmt.Printf("c  等于  %d\n", cc)
}