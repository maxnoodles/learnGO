package main

import "fmt"

func printArr(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printPointArr(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]bool

	fmt.Println(arr1, arr2, arr3, grid)
	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArr(arr1)
	fmt.Println("arr1: ", arr1)
	printArr(arr3)
	fmt.Println("arr1: ", arr3)
	printPointArr(&arr3)
	fmt.Println("arr1: ", arr3)

}
