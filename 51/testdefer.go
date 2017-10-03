package main

import (
	"fmt"
)

func main() {
	testdefer()
	defer testdefer2()
	testdefer3()
}

func testdefer() {
	fmt.Println("I am testdefer1")
}

func testdefer2() {
	fmt.Println("I am testdefer2")
}

func testdefer3() {
	fmt.Println("I am testdefer3")
}
