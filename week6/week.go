package main

import (
	"fmt"
	"time"
)

func getClassDate() {
	today := int(time.Now().Weekday())
	interval := 6 - today
	if interval == 0 {
		interval = 7
	}

	nextClassDate := time.Now().Add(time.Duration(time.Duration(interval*24) * time.Hour))
	fmt.Println(nextClassDate.Format("2006-01-02"))

	for i := 0; i < 3; i++ {
		nextClassDate = nextClassDate.Add(time.Duration(time.Hour * time.Duration(24*7)))
		fmt.Println(nextClassDate.Format("2006-01-02"))
	}
}
