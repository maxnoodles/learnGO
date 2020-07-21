package main

import (
	"fmt"
	"regexp"
)

const text = `My email is noodles@gmail.com.cn
email1 is abc@abc.org
email2 is  vcx@dfa.com`

func main() {
	re := regexp.MustCompile(`\w+@[\w.]+\.\w+`)
	match := re.FindAllString(text, -1)
	fmt.Println(match)
}
