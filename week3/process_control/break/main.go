package main

//func add(a int) int {
//	return a + 10
//}
//
//func switch_expression() {
//	var a int = 5
//	switch add(a) {
//	case 15:
//		fmt.Println("right")
//	default:
//		fmt.Println("wrong")
//	}
//
//	const B = 15
//	switch B {
//	case add(a):
//		fmt.Println("Bright")
//	default:
//		fmt.Println("Bwrong")
//	}
//
//	fmt.Printf("%v", add(a))
//}

func switch_type() {
	//var num interface{} = 1.5
	//switch num.(type) {
	//case int:
	//	fmt.Println("int")
	//case float32:
	//	fmt.Println("float32")
	//case float64:
	//	fmt.Println("float64")
	//case byte:
	//	fmt.Println("byte")
	//case string:
	//	fmt.Println("string")
	//default:
	//	fmt.Println("neither")
	//}

	//var num interface{} = "1.5"
	//switch value := num.(type) {
	//case int:
	//	fmt.Println("number is int %d\n", value)
	//case float64:
	//	fmt.Println("number is float64 %f\n", value)
	//case byte, string:
	//	fmt.Println("number is interface %v\n", value)
	//default:
	//	fmt.Println("neither")
	//}

	//switch num.(type) {
	//case int:
	//	value := num.(int)
	//	fmt.Printf("number is int %d\n", value)
	//case float64:
	//	value := num.(float64)
	//	fmt.Printf("number is float64 %f\n", value)
	//case byte:
	//	value := num.(byte)
	//	fmt.Printf("number is byte %d\n", value)
	//default:
	//	fmt.Println("neither")
	//}

	//arr := []int{1, 2, 3, 4, 5}
	//for i := 0; i < len(arr); i++ {
	//	fmt.Printf("%d:%d\n", i, arr[i])
	//}
}

func main() {
	//switch_expression()
	switch_type()
}
