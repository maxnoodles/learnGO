package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "ccmouse",
		"course": "golang",
		"site":   "imooc",
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m["name"])

	if v, ok := m["abc"]; ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("key not existed")
	}

	fmt.Println("Delete values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
