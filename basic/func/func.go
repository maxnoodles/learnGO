package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(op string, a, b int) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pointSwap(a, b *int) {
	*a, *b = *b, *a
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	if ret, err := eval("X", 3, 4); err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(ret)
	}
	fmt.Println(eval("/", 12, 4))
	fmt.Println(div(12, 4))
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	a, b := 3, 4
	pointSwap(&a, &b)
	fmt.Println(a, b)
	fmt.Println(swap(a, b))
}
