package test

import (
	"fmt"
)

func Reverse(input string) string {

	s := []rune(input)

	fmt.Println("reverse ", s)

	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		t := s[i]
		s[i] = s[j]
		s[j] = t
	}

	return string(s)
}
