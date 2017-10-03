package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
	no   int
}

func main() {
	var xiaoming student
	xiaoming = student{age: 100}
	fmt.Println(xiaoming)
	fmt.Println("name=", xiaoming.name, "age=", xiaoming.age, "no=", xiaoming.no)
}

// 注意， 重新复习了一下struct， 了解到可以给部分的变量赋值，还有个名字叫做符文
// 写代码时候，不对的地方，只有student{age：100}的地方，使用了括号（），因为想到可能是new出来一个新的，类似构造函数所以用括号啊，结果是{}
