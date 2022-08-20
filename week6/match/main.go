package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//source := rand.NewSource(time.Now().UnixMilli())
	//rander := rand.New(source)
	//for i := 0; i < 10; i++ {
	//	fmt.Println(rander.Intn(500))
	//}
	//fmt.Println(rand.Intn(500))

	arr := rand.Perm(100)
	fmt.Println(arr)
}
