package main

import "fmt"

func slice_init() {
	//s2d := [][]int{
	//	{1},
	//	{3, 5, 7},
	//}
	//fmt.Println(len(s2d))
	//fmt.Println(len(s2d[0]))
	//fmt.Println(len(s2d[1]))

	var s []int
	fmt.Printf("len %d,cap %d\n", len(s), cap(s))
	s = []int{}
	fmt.Printf("len %d,cap %d\n", len(s), cap(s))
	s = make([]int, 3, 5)
	fmt.Printf("len %d,cap %d\n", len(s), cap(s))
	s = make([]int, 3)
	fmt.Printf("len %d,cap %d\n", len(s), cap(s))
}

func slice_append() {

	// 指定长度为3，即初始化为 0 0 0，当 append 时，变成 0 0 0 $number，除非指定长度为 0,当 append 时，才会变成 $number
	s := make([]int, 3, 5)

	//s := []int{2, 4, 5, 6}
	fmt.Println(s)
	fmt.Printf("len %d,cap %d\n", len(s), cap(s))
	s = append(s, 100)
	fmt.Println(s)
	fmt.Printf("len %d,cap %d\n", len(s), cap(s))
}

func coef_cap() {
	s := make([]int, 0, 5)
	prevCap := cap(s)
	for i := 0; i < 10000; i++ {
		s = append(s, 0)
		currCap := cap(s)
		if currCap > prevCap {
			fmt.Printf("cap %d ----> %d\n", prevCap, currCap)
			prevCap = currCap
		}
	}
}

func sub_slice() {
	arr := make([]int, 3, 5)
	fmt.Printf("len %d,cap %d\n", len(arr), cap(arr))
	fmt.Println(arr)
	crr := arr[0:2]
	crr[1] = 8
	fmt.Println(arr[1])
	crr = append(crr, 9)
	fmt.Println(arr[2])
	fmt.Println(arr)
	crr = append(crr, 9)
	crr = append(crr, 9)
	fmt.Println(arr)
	fmt.Println(crr)
	fmt.Printf("%p\n %p\n", &arr[0], &crr[0])

	// 子切片超过母切片长度，发生内存分离
	crr = append(crr, 9)
	crr = append(crr, 9)
	crr[1] = 10
	fmt.Println(arr[1])
	fmt.Printf("%p\n %p\n", &arr[0], &crr[0])
}

func update_slice(arr []int) {
	arr[0] = 100
}

func main() {
	//slice_init()
	//slice_append()
	//coef_cap()
	//sub_slice()
	crr := []int{3, 4, 5}
	update_slice(crr)
	fmt.Println(crr[0])
}
