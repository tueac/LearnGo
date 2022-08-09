package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fuzhi() {
	s1 := "dsfdsfds fdsfds dfaf1 232123 k中城"
	s2 := "fdsf '\"fdsfsf\tfd3424\n \\ \\中国 人"
	s3 := `在地erwerew "  fdfd 'g

\' 中"
`
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func StringTest() {
	s := "abc123cfgabc中国人abc"
	fmt.Println(len(s))
	fmt.Println(strings.Split(s, "abc"))
	fmt.Println(strings.Contains(s, "abc"))
	fmt.Println(strings.Contains(s, "abcd"))
	fmt.Println(strings.Index(s, "abc"))
	fmt.Println(strings.LastIndex(s, "abc"))
}

func Context() {
	s1 := "abc"
	s2 := "123"
	s3 := "中国人"

	s4 := s1 + s2 + s3
	fmt.Println(s4)
	s5 := fmt.Sprintf("%s%s%s", s1, s2, s3)
	fmt.Println(s5)
	arr6 := []string{s1, s2, s3}
	s6 := strings.Join(arr6, "")
	fmt.Println(s6)
}

func StringDi() {
	s := "w工"
	arr := []byte(s) //强制类型转换，[]byte 为需要转的类型
	brr := []rune(s)
	fmt.Println(arr, brr)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])

	fmt.Println("--------")
	for _, ele := range s {
		fmt.Printf("%c\n", ele)
	}
	fmt.Println("--------")
}

func cast() {
	var i int = 100
	var b byte = byte(i)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var s string = strconv.Itoa(i)
	fmt.Printf("%s\n%T\n", s, s)
	var err error
	var q int
	q, err = strconv.Atoi(s)
	if err == nil {
		fmt.Printf("%d\n", q)
	}
	var f float32 = 3.14
	s = strconv.FormatFloat(float64(f), 'f', 4, 64)
	fmt.Printf("%s\n", s)
}

func main() {
	//fuzhi()
	//StringTest()
	//Context()
	//StringDi()
	cast()
}
