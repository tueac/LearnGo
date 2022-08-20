package main

import (
	"fmt"
	"time"
)

type user struct {
	id         int
	score      float32
	enrollment time.Time
	name, addr string
}

func (u user) hello(abc string) {
	fmt.Println("hi " + abc + ", my name is" + u.name)
}

func (_ user) he1(abc string) {
	fmt.Printf("hi " + abc + ", my name is")
}

func (user) he2(abc string) {
	fmt.Printf("hi " + abc + ", my name is")
}

func (abcd user) do(a float32) {
	fmt.Printf("%.1f %.1f\n", a, abcd.score)
}

var stu struct {
	Abb string
}

func goooo() {
	var u1 user
	u1 = user{id: 34, addr: "wh"}

	var u2 *user
	u2 = &user{id: 55, addr: "hb"}
	fmt.Println(u1.addr, u2.addr)
}

func copy() {
	var u1 = user{name: "wh"}
	var u2 user
	u2 = u1
	u2.name = "hb"
	fmt.Println(u1.name)
	fmt.Println(u2.name)
}

func main() {
	//var arr []user
	//arr = make([]user, 5)
	//fmt.Println(arr[4].id)

	//var u user
	////	u = user{id: 12, name: "mili", score: 150, enrollment: time.Now()}
	//u = user{12, 150, time.Now(), "mili", "wuhan"}
	//fmt.Println(u.score)
	//fmt.Println(u.enrollment)

	//a := struct {
	//	Name string
	//}{Name: "match"}
	//fmt.Println(a.Name)

	//goooo()

	copy()
}
