package main

import (
	"fmt"
)

func main() {
	var _ int
	_ = 200
	_ = 100
	var _ byte
	_ = "string"

	var _1, _a, __ int
	__ = 300

	var a int
	a = 100
	//	var a byte // redeclared error

	fmt.Println(_1, _a, __, a)
}
