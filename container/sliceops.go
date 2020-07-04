package main

import "fmt"

func printSlice(slice []int) {
	fmt.Printf("len=%d, caps=%d\n",
		len(slice), cap(slice))
}

func main() {
	var arr []int
	for i := 0; i < 100; i++ {
		arr = append(arr, i)
		printSlice(arr)
	}
	fmt.Println(arr)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	copy(s2, arr)
	fmt.Println("Copy slice", s2)
	printSlice(s2)

	fmt.Println("Delete elements")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
	fmt.Println(s2)

	fmt.Println(" Poping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front, s2)

	fmt.Println(" Poping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail, s2)
}
