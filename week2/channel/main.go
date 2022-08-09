package main

import (
	"fmt"
)

func chan1() {
	var ch chan int
	//fmt.Printf("ch is nil %t\n", ch == nil)
	//fmt.Printf("len of ch is %d\n", len(ch))

	ch = make(chan int, 10)
	//fmt.Printf("len os ch is %d\n", len(ch))
	//fmt.Printf("cap of ch is %d\n", cap(ch))

	// ch <- 3
	// ch <- 3
	//fmt.Printf("len of ch is %d\n", len(ch))
	//fmt.Printf("cap of ch is %d\n", cap(ch))

	for i := 0; i < 10; i++ {
		ch <- 3
	}

	fmt.Printf("len of ch is %d\n", len(ch))
	fmt.Printf("cap os ch is %d\n", cap(ch))

	a := <-ch
	<-ch
	fmt.Println(a)
	ch <- 3

	fmt.Printf("len of ch is %d\n", len(ch))
	fmt.Printf("cap of ch is %d\n", cap(ch))
}

func for_chan() {
	ch2 := make(chan int, 10)

	ch2 <- 2
	for i := 0; i < 9; i++ {
		ch2 <- 2
	}

	fmt.Printf("len of ch2 is %d\n", len(ch2))
	fmt.Printf("cap of ch2 is %d\n", cap(ch2))

	close(ch2)
	// 等价于下者，len 将为 0
	for i := 0; i < len(ch2); i++ {
		ele := <-ch2
		fmt.Println(ele)
	}

	// 等价于上者，len 将为 0
	// for ele := range ch2 {
	// 	fmt.Println(ele)
	// }

	fmt.Println("------------")
	fmt.Printf("len of ch2 is %d\n", len(ch2))
	fmt.Printf("cap of ch2 is %d\n", cap(ch2))
}

func update_slice(arr []int) {
	// var arr3 chan int
	// arr3 = make(chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	arr3 <- 10
	// }
	// fmt.Printf("len of arr3 is %d\n", len(arr3))
	// fmt.Printf("cap of arr3 is %d\n", cap(arr3))

	arr[0] = 100
}

func main() {
	// chan1()
	//for_chan()
	crr := []int{3, 4, 5}
	update_slice(crr)
	fmt.Println(crr[0])
}
