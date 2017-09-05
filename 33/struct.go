package main

import "fmt"

func main() {
	
type   person struct {
	firstname string;
	lastname string;
	age int
	}

var gao person

hao:=person {"xiao","mei",30}

fmt.Println(hao.firstname, " " , hao.lastname, " ", hao.age)

gao = hao

fmt.Println(gao.firstname, " " , gao.lastname, " ", gao.age)

var qiao *person
qiao = &gao

fmt.Println(qiao.firstname, " " ,qiao.lastname, " ", qiao.age)

}
