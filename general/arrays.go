package main

import "fmt"

func Arrays() {
	// var arr [5]int
	arr := [3]int{1, 2, 3}
	// arr[0] = 100
	// arr[3] = 900
	fmt.Println("array:", arr)
	fmt.Println(len(arr))
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("sum:", sum)
	arr2D := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("2d:", arr2D)
}
