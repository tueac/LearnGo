package main

import "fmt"

func basic() {
	var arr1 [5]int = [5]int{}
	var arr2 = [5]int{}
	var arr3 = [5]int{1, 2, 4}
	var arr4 = [5]int{1: 12, 0: 10, 2: 88}
	var arr5 = [...]int{1, 2, 4, 1000}
	var arr6 = [...]struct {
		name string
		age  int
	}{{"tom", 18}, {"jim", 20}}
	var arr7 = [5]string{"d", "1", "我"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(len(arr4))
	fmt.Println(arr4[2], arr4[4])
	//arr4[4] = 18
	arr4[len(arr4)-1] = 18
	fmt.Println(arr4[2], arr4[4])
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr6[1])
	fmt.Println(arr7)
	fmt.Println(arr7[2], arr7[4], arr7[0])
}

func Test() {
	var arr1 [3]bool = [3]bool{}
	fmt.Println(arr1)
	arr1 = [3]bool{false, true, false}
	fmt.Println(arr1)

	var arr2 = [3]float64{1, 2.01, 3}
	fmt.Println(arr2)

	var arr3 = [...]string{"te", "li", ""}
	fmt.Println(len(arr3), cap(arr3), arr3)

	var arr4 = [3]float64{2: 3.14}
	fmt.Println(arr4)
}

func arrPoint(arr *[5]int) {
	fmt.Println(arr[0])
	fmt.Printf("%p\n", arr)
	arr[0] += 10
	fmt.Println(arr[0])
}

func array2d() {
	var arr6 = [5][3]int{{1, 2}, {3}, {6, 3, 8}}
	var arr7 = [...][3]int{{1, 2}, {1, 8}}
	fmt.Println(len(arr6))
	fmt.Println(len(arr7))

	//for i, row := range arr6 {
	//	fmt.Printf("%T\n", row)
	//	for j, n := range row {
	//		fmt.Printf("%d %d %d\n", i, j, n)
	//	}
	//}

	for i := 0; i < len(arr6); i++ {
		for j := 0; j < len(arr6[i]); j++ {
			fmt.Printf("%d %d %d\n", i, j, arr6[i][j])
		}
	}
}

func main() {
	//server()

	//var crr [5]int = [5]int{1, 2, 6, 8, 10}
	//fmt.Printf("%p\n", &crr)
	//arrPoint(&crr)
	//fmt.Println(crr[0])
	//
	//fmt.Println(len(crr))
	//fmt.Println(cap(crr))

	//array2d()
	//basic()
	Test()
}
