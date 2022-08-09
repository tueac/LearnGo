package main

import (
	"fmt"
)

func foo(i interface{}) {
	if v, ok := i.(int); ok {
		fmt.Printf("i is int %d\n", v)
	} else if v, ok := i.(float32); ok {
		fmt.Printf("i is float32 %f\n", v)
	} else if v, ok := i.(byte); ok {
		fmt.Printf("i is byte %b\n", v)
	} else {
		fmt.Printf("other type")
	}
}

func f2(i interface{}) {
	switch i.(type) {
	case int:
		v := i.(int)
		fmt.Printf("i is int %d\n", v)
	case float32:
		v := i.(float32)
		fmt.Printf("i is float32 %f\n", v)
	case byte:
		v := i.(byte)
		fmt.Printf("i is byte %b\n", v)
	default:
		fmt.Printf("other type")
	}
}

func main() {
	var i int = 8
	foo(i)
	f2(i)

	var f float32
	foo(f)
	f2(f)
}
