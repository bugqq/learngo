package main

import (
	"errors"
	"fmt"
)

type CalAge struct {
	now   float64
	birth float64
}

func (age CalAge) Error() string {
	message := "now = %f, birth = %f, now < birth is error\n"
	return fmt.Sprintf(message, age.now, age.birth)
}

func checkage(age CalAge) (bool, string) {
	if age.now < age.birth {
		return false, age.Error()
	} else {
		return true, ""
	}
}

func main() {
	age1 := CalAge{100.0, 200.0}
	fmt.Println(age1.Error())
	v, e := checkage(age1)
	fmt.Println(v, e)
	if e == "" {
		fmt.Println("error is nil")
	}
	age1 = CalAge{200.0, 100.0}
	v, e = checkage(age1)
	fmt.Println(v, e)
	if e == "" {
		fmt.Println("error is nil")
	}

	// new error
	err := errors.New("work later and back to home until 10pm, 2017/09/11")
	fmt.Println(err)
}
