package main

import "fmt"

func main() {
	var names [2] **int
	var value = []int{100,200} 
	for i:=0;i < 2;i ++ {
		temp:=&value[i]
		names[i] = &temp
	}
	fmt.Printf("%d\n", **names[0])
	fmt.Printf("%d\n", **names[1])
}