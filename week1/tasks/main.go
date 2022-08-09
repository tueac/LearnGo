package main

import (
	"fmt"
	"math"
	"strings"
)

//输出一个int32对应的二进制表示
func BinaryFormat(n int32) string {
	a := uint32(n)               //有符号转为无符号
	sb := strings.Builder{}      //把多个短字符串拼接成一个长字符串
	c := uint32(math.Pow(2, 31)) //最高位是1，其他位全是0

	for i := 0; i < 32; i++ {
		if a&c != 0 {
			sb.WriteString("1") //往sb后面拼接“1”
		} else {
			sb.WriteString("0") //往sb后面拼接“0”
		}
		c >>= 1
	}
	return sb.String() //返回长度为32的字符串
}

func main() {
	fmt.Println(BinaryFormat(0))
	fmt.Println(BinaryFormat(-32))
	fmt.Println(BinaryFormat(128))
	fmt.Println(BinaryFormat(100000000))
}
