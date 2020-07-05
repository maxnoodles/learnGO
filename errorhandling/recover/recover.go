package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occured", err)
		} else {
			panic(fmt.Sprintf("I don't not what to do: %v", r))
		}
	}()

	//panic(errors.New("this is a error"))
	//a, b := 5, 0
	//fmt.Println(a / b)
	panic(123)
}

func main() {
	tryRecover()
}
