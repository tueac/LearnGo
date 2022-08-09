package main

import "math/rand"

func forArrary() {
	array := [20]int{}
	for i := 0; i < 20; i++ {
		array[i] = rand.Int()
	}

	sum := 0
	product := 1

	var i int = 0
	// 求前 10 个元素的和
	for ; i < 10; i++ {
		sum += array[i]
	}

	// 求后 10 个元素的积
	for ; i < 20; i++ {
		product *= array[i]
	}
}
