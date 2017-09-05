package main

import "fmt"

// test for append and copy
func main() {

	range2 := make([]int, 10, 15)
	// append(range2, 100) // this will be error, evaluated but not used
	range2 = append(range2, 100)
	fmt.Printf("%v, %d, %d\n", range2, len(range2), cap(range2))

	range2 = append(range2, 100, 200, 300, 400)
	fmt.Printf("%v, %d, %d\n", range2, len(range2), cap(range2))

	range2 = append(range2, 100, 100)
	fmt.Printf("%v, %d, %d\n", range2, len(range2), cap(range2))

	subrange2 := []int{}
	copy(subrange2, range2)
	fmt.Printf("%v, %d, %d\n", subrange2, len(subrange2), cap(subrange2))

	subrange3 := make([]int, 5, 10)
	copy(subrange3, range2)
	fmt.Printf("%v, %d, %d\n", subrange3, len(subrange3), cap(subrange3))

}
