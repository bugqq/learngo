package main

import "fmt"

type learn interface {
	learn()
}

type language struct {
	name string
}

type physics struct {
	name string
}

func (lang language) learn() {
	fmt.Println("learn go")
	fmt.Println(lang.name)
}
func (phy physics) learn() {
	fmt.Println("learn physics")
	fmt.Println(phy.name)
}

func main() {
	var java language
	java.name = "java"
	java.learn()

	earch := physics{"\n, a cource about earch"}
	earch.learn()
}
