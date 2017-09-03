package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var c int = 400

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	* &a 指向 a 变量的地址
	* b 指向 b 变量的地址
	* c 指向 c 变量的地址地址
	 */
	addressofc := &c
	fmt.Printf("交换前 b 的地址 : %x\n", &b)
	fmt.Printf("交换前 c 的值 : %d\n", c)
	swap(&a, b, &addressofc)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
	fmt.Printf("交换后 c 的值 : %d\n", c)
	fmt.Printf("交换后 b 的地址 : %x\n", &b)
}

func swap(x *int, y int, z **int) {
	y = 300
	*x = y
	fmt.Printf("交换中 b 的地址 : %x\n", &y)
	**z = *x
}
