package main

import (
	"fmt"
	"time"
)

func time1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("20060102150405"))
}

func time2() {
	now := time.Now().Format("20060102150405")
	fmt.Println(now)
}

func time3() {
	TimeFormat := "2006-01-02 15:04:05"
	now := time.Now()
	fmt.Println(now.Format(TimeFormat))
}

func time4() {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	TimeFormat := "2006-01-02 15:04:05"
	now := time.Now()
	fmt.Println(now.Unix(), now.UnixMilli(), now.UnixMicro(), now.UnixNano())
	nowStr := now.Format(TimeFormat)
	fmt.Println(nowStr)

	if nt, err := time.Parse(TimeFormat, nowStr); err == nil { // err == nil，err 没有出错
		fmt.Println(nt.Year(), int(nt.Month()), nt.Day(), nt.Hour(), nt.Minute(), nt.Second())
	} else {
		fmt.Println(err)
	}

	// 避免远程主机时间非东八区时，显示的时间与预期不一致
	if ntr, err := time.ParseInLocation(TimeFormat, nowStr, loc); err == nil {
		fmt.Println(ntr.Year(), int(ntr.Month()), ntr.Day(), ntr.Hour(), ntr.Minute(), ntr.Second(), ntr.YearDay(), int(ntr.Weekday()))
	} else {
		fmt.Println(err)
	}
}

func time5() {
	Tm := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now().Unix())
	<-Tm.C
	fmt.Println(time.Now().Unix())
	Tm.Stop()

	fmt.Println(time.Now().Unix())
	<-time.After(3 * time.Second)
	fmt.Println(time.Now().Unix())

	fmt.Println(time.Now().Unix())
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now().Unix())
}

func time6() {
	Tk := time.NewTicker(500 * time.Millisecond)
	for i := 0; i < 10; i++ {
		<-Tk.C
		fmt.Println(time.Now().UnixMilli())
	}
	Tk.Stop()
}
func main() {
	//time1()
	//time2()
	//time3()
	//time4()
	//time5()
	time6()
}
