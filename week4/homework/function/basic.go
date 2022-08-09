package main

import "fmt"

func argf(a, d int, b, c string) (int, string) {
	// a b c d 称为形参
	// 形参是实参的拷贝
	a += d
	b += c
	e := a
	return e, b
}

func useAppend() {
	arr := []int{1, 2, 3}
	brr := append(arr, 5, 6, 7, 8)
	fmt.Println(arr)
	fmt.Println(brr)

	crr := append(arr, brr...) // 在切片后面添加 ...，把切片转成了不定长参数
	fmt.Println(crr)
}

func defer_exec_time() (i int) {
	i = 9
	defer func() { // 放在匿名函数中，打印 5
		fmt.Printf("i=%d\n", i)
	}()
	defer fmt.Printf("i=%d\n", i) // 没有放在匿名函数中，变量注册 defer 时被拷贝
	return 5
}

func main() {
	//useAppend()
	defer_exec_time()

	x, y := 3, 6
	m, n := "123", "abc"     // x y m n 称为实参
	w, q := argf(x, y, m, n) // w q 是 e b 的拷贝
	fmt.Printf("%p %p\n", &w, &q)
	fmt.Println(w, q)
}
