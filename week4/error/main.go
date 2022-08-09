package main

import (
	"fmt"
	"time"
)

type PathError struct {
	path    string
	op      string
	opTime  time.Time
	message string
}

func (err PathError) Error() string {
	return err.opTime.Format("20060102") + ":" + err.op + " " + err.path + err.message
}

func deletePath(path string) error {
	return PathError{path: path, op: "delete", opTime: time.Now(), message: "目录不存在"}
}

func main() {
	path := "/usr/not/use"
	if err := deletePath(path); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("目录%s删除成功\n", path)
	}

	foo()
}
