package main

import "fmt"

func main() {

	x := recursion(5)
	fmt.Println(x)
}

func recursion(x int) int {
	if x <= 1 {
		x = 1
	} else {
		x = x * recursion(x-1)
	}
	return x
}
