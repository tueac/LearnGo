package main

import "fmt"

func main() {
	var m map[string]int
	//fmt.Println(len(m))

	m = make(map[string]int)
	//fmt.Println(len(m))

	// cap 为 10，map 不支持 cap 函数
	m = make(map[string]int, 10)
	//fmt.Println(len(m))

	m = map[string]int{"A": 3, "B": 2}
	//fmt.Println(len(m))

	// 增加一个 key，value
	m["D"] = 18
	//fmt.Println(len(m))

	// 删除一个 key，value
	delete(m, "A")
	//fmt.Println(len(m))

	// 删除一个 不存在 key，value
	delete(m, "张三")
	//fmt.Println(len(m))

	//value := m["Q"]
	//fmt.Println(value)
	//value = m["A"]
	//fmt.Println(value)

	key := "BB"
	v, ok := m[key]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Printf("key %s 不存在\n", key)
	}

	// for key, value := range m {
	// 	fmt.Printf("key %s, value %d\n", key, value)
	// }
}
