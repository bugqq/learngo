package main

import "fmt"

func main() {

	range1 := []int{1, 2, 3}
	fmt.Printf("%v\n", range1)

	range2 := make([]int, 10, 15)
	fmt.Printf("%v\n", range2)

	var range3 []int
	var range4 [3]int = [3]int{1, 2, 3}
	range3 = range4[1:3]
	fmt.Printf("%v\n", range3)

	len := len(range3)
	cap := cap(range3)
	fmt.Printf("len=%d,cap=%d\n", len, cap)

	// will occur error, because the index out of rage for length 10
	//range2[10] = 100
	//fmt.Println(range2[10])

	range2 = range3[:1]
	fmt.Printf("%v\n", range2)

}
