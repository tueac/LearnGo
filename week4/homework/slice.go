package main

import "fmt"

func d2slice() {
	array := [5][4]float64{}
	scl := make([]float64, 0, 20)
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			scl = append(scl, array[i][j])
		}
	}
}

func appendSclice() {
	arr := make([]int, 3, 4)
	brr := append(arr, 1)
	arr[0] = 8
	fmt.Println("brr is", brr[0])
	fmt.Printf("%p", &brr)

	brr = append(brr, 6)
	brr[0] = 9
	fmt.Println("arr is", arr[0])
	fmt.Printf("%p", &arr)
	fmt.Println("brr is append", brr[0])
	fmt.Printf("%p\n", &brr)
}
