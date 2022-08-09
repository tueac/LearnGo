package main

import (
	"fmt"
	"learn-go/week1/util"
	"runtime"
	"strconv"
)

func airthmetic() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = a + b
	var d float32 = a - b
	var e float32 = a * b
	var f float32 = a / b
	fmt.Printf("a=%.4f\nb=%.4f\nc=%.4f\nd=%.4f\ne=%.4f\nf=%.4f\n", a, b, c, d, e, f)
}

func relational() {
	var a float32 = 8
	var b float32 = 8
	fmt.Printf("a等于b吗？\n%t\n\n", a == b)
	fmt.Printf("a大于b吗？\n%t\n\n", a > b)
	fmt.Printf("a小于b吗？\n%t\n\n", a < b)
	fmt.Printf("a不等于b吗？\n%t\n\n", a != b)
	fmt.Printf("a小于或等于b吗？\n%t\n\n", a <= b)
	fmt.Printf("a大于或等于b吗？\n%t\n", a >= b)
}

func logistic() {
	var a float32 = 4
	var b float32 = 2
	var c float32 = 4
	fmt.Printf("a大于b而且b大于c吗？\n%t\n\n", a > b && b > c)
	fmt.Printf("a大于b或者b大于c吗？\n%t\n\n", a > b || b > c)
	fmt.Printf("a大于b是不成立的，对吗？\n%t\n\n", !(a > b))
	fmt.Printf("b大于c是成立的，对吧？\n%t\n", b > c)
}

func bitOp() {
	fmt.Printf("os arch %s, int size %d\n", runtime.GOARCH, strconv.IntSize)
	var a int = 260
	//var b int = -260
	fmt.Printf("a 对应的二进制是是%b\n", a)
	fmt.Printf("a 与 4 对应的二进制值是%b\n", a&4)
	fmt.Printf("a 或 4 对应的二进制值是%b\n", a|4)
	fmt.Printf("a 异或 4 对应的二进制值是%b\n", a^4)
	fmt.Printf("a 按位取反后对应的二进制是%b\n", ^a)
	fmt.Printf("a 右移 2 位后对应的二进制是%b\n", a>>2)
	fmt.Printf("a 左移 2 位后对应的二进制是%b\n", a<<2)
	//fmt.Printf("-b 对应的二进制是%b\n", b)
	//fmt.Printf("b 与 4对应的二进制值是%b\n", b&4)
	//fmt.Printf("b 或 4对应的二进制值是%b\n", b|4)
	//fmt.Printf("b 异或 4对应的二进制值是%b\n", b^4)
	//fmt.Printf("b 按位取反 4对应的二进制值是%b\n", ^b)
	//fmt.Printf("b 右移 2 位对应的二进制值是%b\n", b>>2)
	//fmt.Printf("b 左移 2 位对应的二进制值是%b\n", b<<2)
}

func assignment() {
	var a, b int = 8, 4
	a += b
	fmt.Printf("a = a + b 的值为 %d\n", a)
	a, b = 8, 4
	a -= b
	fmt.Printf("a = a - b 的值为 %d\n", a)
	a, b = 8, 4
	a *= b
	fmt.Printf("a = a * b 的值为 %d\n", a)
	a, b = 8, 4
	a /= b
	fmt.Printf("a = a / b 的值为 %d\n", a)
	a, b = 8, 4
	a %= b
	fmt.Printf("a = a / b 的值为 %d\n", a)
	a, b = 8, 4
	a <<= b
	fmt.Printf("a = a << b 的值为 %d\n", a)
	a, b = 8, 4
	a >>= b
	fmt.Printf("a = a >> b 的值为 %d\n", a)
	a, b = 8, 4
	a &= b
	fmt.Printf("a = a & b 的值为 %d\n", a)
	a, b = 8, 4
	a |= b
	fmt.Printf("a = a | b 的值为 %d\n", a)
	a, b = 8, 4
	a ^= b
	fmt.Printf("a = a ^ b 的值为 %d\n", a)
}

func anD() {
	var a string = "china"
	fmt.Printf("%s\n", a)
	const PI float32 = 5.20

	const (
		// 都是 100，与 a 相等
		aa = 100
		b
		c
		d
		e
	)
	fmt.Printf("%d\n%d\n%d\n%d\n%d\n", a, b, c, d, e)

	const (
		// iota 为 0,依此相加
		f = iota
		g
		h
	)
	fmt.Printf("%d\n%d\n%d\n", f, g, h)

	const (
		i = iota
		j = 30
		k = iota
		l
	)
	fmt.Printf("%d\n%d\n%d\n%d\n", i, j, k, l)

	const (
		m = iota
		n
		_
		p
	)
	fmt.Printf("%d\n%d\n%d\n", m, n, p)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
	)
	fmt.Printf("%d\n%d\n%d\n%d\n", KB, MB, GB, TB)
	fmt.Println(1024 * 1024)
	fmt.Println(1 << 40)

	fmt.Printf("%t\n", 04 == 4.00)
	fmt.Printf("%v\n", .4i)
	fmt.Printf("%U\n%U\n", '罗', '立')
	fmt.Printf("%t\n", '\u7F57' == '罗')
	fmt.Printf("%t\n", '\u7ACB' == '立')
}

func varPack() {
	var c int
	c = util.Add(10, 2)
	fmt.Printf("%d\n", c)

	d := util.Jian("罗", "立")
	fmt.Printf("%s\n", d)

	e := util.Jian("123", "456")
	fmt.Printf("%s\n", e)
}

var ab = 12

func varQ() {
	fmt.Printf("%v   %T\n", ab, ab)
	ab = 10
	fmt.Printf("%v   %T\n", ab, ab)
	{
		ab = 8
		fmt.Printf("%v   %T\n", ab, ab)
	}
}

func main() {
	//airthmetic()
	//relational()
	//logistic()
	//bitOp()
	//assignment()
	//anD()
	//varPack()
	//varQ()
	fmt.Println(12 & 4)
}
