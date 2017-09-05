package main

import "fmt"
	
type   person struct {
	firstname string;
	lastname string;
	age int
	}

func main() {


hao:=person {"xiao","mei",30}

fmt.Println(hao.firstname, " " , hao.lastname, " ", hao.age)

var qiao *person
qiao = &hao
qiao.age = 40

fmt.Println(qiao.firstname, " " ,qiao.lastname, " ", qiao.age)

test_struct_point(qiao)

}

func test_struct_point(pp *person) {
	fmt.Println (pp.firstname , " " , pp.lastname , " " , pp.age)
}
