package main

import "fmt"

func main() {

	index:=1
	for index<10 {
		fmt.Print("hello ")
		index+=1
		break
	}

	var a [10]int 
	a=[10]int{1,2,3,4,5}

	fmt.Print("\n")
	for i,x:= range a {
		fmt.Print(i,  x,"\n")
		continue
		fmt.Print(i,  x,"\n")
	}
	
	gotolabel:
	for i,x:= range a {
		fmt.Print(i,  x,"\n")
		goto gotolabel
		fmt.Print(i,  x,"\n")
	}
	  fmt.Println("world")
}