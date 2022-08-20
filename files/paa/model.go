package pa1

import (
	pb1 "test/pbb"
	uta "test/pbb/abc"
)

type Student struct {
	Name string
	age  int
}

func (s Student) Sqaure() int {
	return pb1.Sqaure(s.age) + uta.Sub(10, 20)
}

type (
	luo struct {
		Name string
	}
	li struct {
		age int
	}
)

func main() {

}
