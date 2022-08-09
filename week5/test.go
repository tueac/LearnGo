package main

import (
	"fmt"
	"time"
)

func argf1(a, b int) {
	a = a + b
	var c = a
	fmt.Println(c)
}

func arg2(a int, b int) {
	a = a + b
}

func arg3(a, b *int) {
	*a = *a + *b
	*b = 888
}

func no_arg() {
	//fmt.Println("欢迎")
}

func return1(a, b int) int {
	a = a + b
	c := a
	return c
}
func return2(a, b int) (c int) {
	a = a + b
	c = a
	return
}

func return3() (int, int) {
	now := time.Now()
	fmt.Println(now)
	return now.Hour(), now.Minute()
}

func main() {
	var x, y = 10, 20
	argf1(x, y)

	return1(x, y)
	return2(x, y)
	return3()
}
