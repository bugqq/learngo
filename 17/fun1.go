package main

import "fmt"

func main() {

	a,b:=100,200
	order(&a,&b)
	  fmt.Println(a, " ", b)

	  
	c,d:=400,300
	order(&c,&d)
	  fmt.Println(c, " ", d)

}

func order(a ,b *int) {
	if  *a<*b {
	 // do nothing
	 } else {
	 t:=*b
	 *b=*a
	 *a=t
	 } 
}
