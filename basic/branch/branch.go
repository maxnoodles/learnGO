package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	const filename = "abc.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

func readFileShortcut() {
	filename := "abc.txt"
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

func grade(source int) string {
	g := ""
	switch {
	case source < 0 || source > 100:
		panic(fmt.Sprintln("this is a wrong source"))
	case source < 60:
		g = "F"
	case source < 80:
		g = "C"
	case source < 90:
		g = "B"
	case source <= 100:
		g = "A"
	}
	return g
}

func main() {
	readFile()
	readFileShortcut()

	fmt.Println(
		grade(0),
		grade(58),
		grade(82),
		grade(100),
	//grade(-5),
	)
}
