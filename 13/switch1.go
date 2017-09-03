package main

import "fmt"

func main() {
	//break确认，输出101
	a,b:=101,200
	switch a {
		case 100:
			fmt.Println("is 100")
		case 101:
			fmt.Println("is 101")
		default:
			fmt.Println("default")
	}

	// default确认，输出default
	switch b {
		case 100:
			fmt.Println("is 100")
		case 101:
			fmt.Println("is 101")
		default:
			fmt.Println("default")
	}

	// 表达式确认， 输出101， 如果是 case 100 和 case a==101的话，会有错误，一个是int，一个是boolean，case的判断值不一致
	switch  {
		case a==100:
			fmt.Println("is 100")
		case a==101:
			fmt.Println("is 101")
		default:
			fmt.Println("default")
	}

	//确认计算表达式，输出200
	switch  {
		case a==100:
			fmt.Println("is 100")
		case a==100 || b == 200:
			fmt.Println("is 200")
		default:
			fmt.Println("default")
	}

	// 确认多值的逗号，输出200
	switch  {
		case a==100:
			fmt.Println("is 100")
		case a==100 ,b == 200:
			fmt.Println("is 200")
		default:
			fmt.Println("default")
	}
}