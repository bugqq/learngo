package main

import "fmt"

type learn interface {
	learn() string
}

type language struct {
	name string
}

type physics struct {
	name string
}

func (lang language) learn() string {
	fmt.Println("learn go")
	fmt.Println(lang.name)
	return "skill"
}
func (phy physics) learn() string {
	fmt.Println("learn physics")
	fmt.Println(phy.name)
	return "ppphhh"
}

func main() {
	var java language
	java.name = "java"
	result := java.learn()

	earch := physics{"\n, a cource about earch"}
	result2 := earch.learn()

	fmt.Println(result, " ", result2)
}
