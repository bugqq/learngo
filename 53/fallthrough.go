package main

import (
	"fmt"
)

func main() {
	var i int
	i = 100
	switch i {
	case 50:
		{
			fmt.Println("50")
		}
	case 100:

		fmt.Println("100")
		fallthrough

	case 200:
		{
			fmt.Println("200")
		}
	case 300:
		{
			fmt.Println("300")
		}
	}
}
