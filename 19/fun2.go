package main

import "fmt"

func main() {
	//初始化
	a:= func()(int,int) {
		return 200,100
	}
	fmt.Println(a, " ")

	//	计算逻辑
	c,d := 100,200
	e:= func(x,y int)(int) {
		return x+y
	}
	fmt.Println(c,d,e)

	// 交换
	f,g := 100,200
		i:= func(x,y int)(int,int) {
		return y,x
	}
	fmt.Println(f,g,i)

	fmt.Println(a())
	fmt.Println(e(c,d))
	fmt.Println(i(f,g))

}

