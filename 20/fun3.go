package main

import "fmt"

func main() {
	//定义
	increment:= func() ( func ()  (int) ) {
		var i int =0
		return func () int {
			i = i +1
			return i
		}
	}
	inc := increment()
	var count int = 0
	count = inc()
	count = inc()
	fmt.Println("count = ", count);

}

