package main

import "fmt"

var 
	(
	a int
	b int
	)

func main() {

a , b = main2()
fmt.Println (a, b)

}

func main2()(int,int) {
return 100,200
}