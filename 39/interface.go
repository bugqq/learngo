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

func (language) learn() {
	fmt.Println("learn go")
}
func (physics) learn() {
	fmt.Println("learn physics")
}

func main() {
	var learn_instance learn
	learn_instance = new(language)
	learn_instance.learn()

	learn_instance = new(physics)
	learn_instance.learn()
}
