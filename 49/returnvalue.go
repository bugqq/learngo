package main

import "fmt"

func main() {
	fmt.Println(testDefaultMultiReturnValue())
	fmt.Println(testDefaultMultiReturnValue2())
}

func testDefaultMultiReturnValue() (x, y, z int) {
	x = 1
	y = 2
	z = 3
	return
}
func testDefaultMultiReturnValue2() (x bool, y, z int) {
	x = true
	y = 2
	z = 3
	return
}
