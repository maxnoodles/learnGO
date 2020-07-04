package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!" // utf-8
	fmt.Println(s, len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	fmt.Println(bytes)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c ", ch)
		bytes = bytes[size:]
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}

}
