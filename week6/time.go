package main

import (
	"fmt"
	"time"
)

func timeFormatConVert() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}

	timeStr := "1998-10-01 08:10:00"
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t.Format("20060102150405"))
	fmt.Println(t.Format("200601021504"))
}
