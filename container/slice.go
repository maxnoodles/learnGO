package main

import "fmt"

func updateSlice(arr []int) {
	arr[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[:]", arr[:])

	s1 := arr[2:]
	fmt.Println("s1=", s1)
	updateSlice(s1)
	fmt.Println("s1=", s1)
	fmt.Println("arr", arr)

	fmt.Println("Extending slice")
	arr[2] = 2
	fmt.Println("arr=", arr)

	s1 = arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))
	fmt.Println(s2[0:4])

	s3 := append(s1, 100)
	fmt.Println(s3)
}
