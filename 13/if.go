package main

import "fmt"

func main() {

	a,b:=101,200
	if a==100 {
		fmt.Println("is 100")
		
	} else {
		fmt.Println("is not 100")
		if b == 200 {
			fmt.Println("is 200")
			
		} else {
			fmt.Println("is not 200")
		}
	}
}