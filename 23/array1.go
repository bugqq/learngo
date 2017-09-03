package main

import "fmt"


func main() {
	name := [...] string {"Han"," " ,"John"}
	fmt.Println("name[0]=",name[0])
	fmt.Println("name[2]=",name[2])

	name[1] = " Peng "
	fmt.Println("name[1]=",name[1])

	temp := name[1]
	fmt.Println("name[1]=",temp)

	var temp2 string
	temp2 = name[1]
	fmt.Println("name[1]=",temp2)
}

