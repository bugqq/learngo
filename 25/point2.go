package main

import "fmt"

func main() {
	var list [3] *int
	color, big,small :=0,1,2
	for i:=0;i<=2;i++ {
		switch i {
			case 0: list[i] = &color
			case 1:list[i] = &big
			case 2:list[i] = &small
		}
	}
	fmt.Printf("%d\n", *list[0])
	fmt.Printf("%d\n", *list[1])
	fmt.Printf("%d\n", *list[2])

}