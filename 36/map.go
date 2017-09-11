package main

import "fmt"

func main() {

	// def
	var carinfo map[string]string

	// init
	carinfo = make(map[string]string)

	// set value
	carinfo["name"] = "eagle"
	carinfo["color"] = "black"
	carinfo["weight"] = "100"

	// get value
	fmt.Println(carinfo)
	fmt.Println(carinfo["name"])

	// loop
	for k, v := range carinfo {
		fmt.Println(k, " ", v)
	}

	// contain
	k, v := carinfo["belong"]
	fmt.Println(k, " ", v)

	// delete
	delete(carinfo, "name")
	fmt.Println("after delete = ", carinfo["name"])

	fmt.Println("end")

	// size
	fmt.Println(len(carinfo))

}
