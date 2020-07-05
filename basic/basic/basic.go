package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 5
	bb = 6
	cc = 7
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d  %q\n", a, s)
}

func variableInit() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "abc"
	b = 5
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "abc"
	b = 6
	fmt.Println(a, b, c, d)
}

func euler() {
	fmt.Printf("%.3f\n\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	a, b := 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a int, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func emun() {
	const (
		cpp = iota
		java
		python
		goland
		js
	)
	fmt.Println(cpp, java, python, goland, js)

	const (
		b = 1 << (10 * iota)
		mb
		kb
		gb
		tb
	)
	fmt.Println(b, mb, kb, gb, tb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInit()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
	euler()
	triangle()
	consts()
	emun()
}
