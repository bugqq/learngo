package main

import "fmt"

func main() {

	// test range
	for key, value := range "go" {
		fmt.Println(key, "=", value)
	}

	for key, value := range []int{1, 2, 3} {
		fmt.Println(key, "=", value)
	}

	for _, value := range []string{"hello", "world", "!"} {
		fmt.Println(value)
	}

	for key, value := range map[string]string{"name": "shame on me", "sex": "undefined"} {
		fmt.Println(key, "-", value)
	}
}
