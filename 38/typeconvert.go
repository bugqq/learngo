package main

import "fmt"

func main() {

	x := recursion(5)
	fmt.Println(float32(x))
	fmt.Printf("%f\n", float32(x))

	var f float32
	f = float32(x)
	fmt.Printf("%f\n", f)
}

func recursion(x int) int {
	if x <= 1 {
		x = 1
	} else {
		x = x * recursion(x-1)
	}
	return x
}
