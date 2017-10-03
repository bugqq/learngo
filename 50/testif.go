package main

import "fmt"

func main() {

	testif(0)
	testif(1)
	testif2(0, 1)
	testif2(1, 2)
}

func testif(studentno int) {
	if studentno == 0 {
		fmt.Println("no is ", studentno)
	} else {
		fmt.Println("no is not zero is ", studentno)
	}
}

func testif2(studentno int, studentno2 int) {
	if allno := studentno + studentno2; allno < 2 {
		fmt.Println("allno < 2 ", studentno, studentno2, allno)
	} else {
		fmt.Println("allno >= 2  ", studentno, studentno2, allno)
	}
}

// 注意点
// 1.不能是两个testif方法啊，即使参数不一样也不行，所以一个改成了testif2
// 2.if 条件的部分，用的是分号，而不是逗号，如何记住呢？
