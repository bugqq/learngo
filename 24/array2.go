package main

import "fmt"


func main() {
	count := [...] int {100,200}
	sum:=sum(count)
	fmt.Println("sum=",sum)

	count2 := [] int {100,200}
	sum2:=sum2(count2)
	fmt.Println("sum2=",sum2)
}

func sum(aArray[2] int) int {
	return aArray[0] + aArray[1]
}


func sum2(aArray[] int) int {
	return aArray[0] + aArray[1]
}