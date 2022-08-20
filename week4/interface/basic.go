package main

import "fmt"

func argf(a int, b, c string, d int) (int, string) {
	a += d
	b += c
	e := a
	fmt.Printf("%s\n%s\n", b, c)
	//fmt.Printf("%p %p\n", &e, &c)
	return e, b
}

func userAppend() {
	arr := []int{1, 2, 3}
	brr := append(arr, 5, 5623, 153256, 55)
	fmt.Println(brr)

	crr := append(arr, brr...)
	fmt.Println(crr)

	sliceT1 := append([]byte("luoli"), ("test1")...)
	sliceT2 := append([]rune("hello"), []rune("word")...)
	fmt.Println(sliceT1, sliceT2)
}

func Fibonacei1(n int) (c int) {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacei1(n-1) + Fibonacei1(n-2)
}

func main() {
	//x, y := 3, 6
	//m, n := "abc", "def"
	//w, q := argf(x, m, n, y)
	//fmt.Println(w, q)
	//fmt.Printf("%p %p\n", &w, &q)
	//userAppend()
	Fibonacei1(2)
}
