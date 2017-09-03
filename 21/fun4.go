package main

import "fmt"

//定义
type person struct {
	age int
}

func main() {
	var xiaohong person
	xiaohong.age = 99

	age_add := xiaohong.age_add()
	fmt.Println("age_add = ", age_add);
	
	var count int = 0
	count = age_add()
	count = age_add()
	fmt.Println("count = ", count);

	for i:=0;i<3;i++ {
		count = age_add()
	}
	fmt.Println("count = ", count);
	

}

func (oneperson person) age_add() (func () (int)) {
	var i int =0
	return func () int {
		i ++
		return oneperson.age + i
	}
}

