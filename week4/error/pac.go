package main

import "fmt"

func foo() {
	defer fmt.Println("1111")
	a, b := 3, 7
	//b = 0
	_ = a / b
	fmt.Println("ok1")
	arr := []int{1, 3, 7}
	defer fmt.Println("2222")
	fmt.Println(arr[21])
	defer fmt.Println("3333")
	fmt.Println("ok2")
}
