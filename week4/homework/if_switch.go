package main

func jijie(month int) string {
	if month < 1 || month > 12 {
		return ""
	}
	if month <= 3 {
		return "春"
	} else if month <= 6 {
		return "夏"
	} else if month <= 9 {
		return "秋"
	} else {
		return "冬"
	}
}

func jijie2(month int) string {
	switch month {
	case 1, 2, 3:
		return "春"
	case 4, 5, 6:
		return "夏"
	case 7, 8, 9:
		return "秋"
	case 10, 11, 12:
		return "冬"
	default: // 非法有份
		return ""
	}
}

// func main() {
// 	fmt.Printf("%s\n", jijie(8))
// 	fmt.Printf("%s\n", jijie2(12))
// }
