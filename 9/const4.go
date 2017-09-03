package main

import "fmt"

func main() {

	const (
		a = iota
		b 
		c  
		d = 100
		e
		f = iota
	)
	fmt.Println (a, b, c,d,e,f)

	
	const (
		a2 = 100
		b2
		c2 = iota
	)
	fmt.Println (a2, b2, c2)

	const (
		i=1<<iota
		j=3<<iota
		k
		l
	)


	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)

}