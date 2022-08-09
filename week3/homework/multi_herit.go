package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Height float32
}

type C struct {
	Sex string
}

type D struct {
	A // C 继承了 A, 匿名才是继承
	B
	c C // 通过 C 去使用 Sex,非继承
}

func (d D) hello1() {
	d.Name = "123"
}

func (d *D) hello2() {
	d.Name = "123"
}
func he() {
	d := D{}
	fmt.Println(d.Name, d.Height, d.c.Sex)
}

func callHello() {
	d := D{}
	d.Name = "abc"

	d.hello1()
	fmt.Println(d.Name) // abc 使用的拷贝

	d.hello2()
	fmt.Println(d.Name) // 123 使用的是指针
}
