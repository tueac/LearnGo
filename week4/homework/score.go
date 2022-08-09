package main

type Student struct {
	Name   string
	Yuwen  float64
	Shuxue float64
	Yingyu float64
	Avg    float64
}

type Class struct {
	students  []Student
	AvgYuwen  float64
	AvgShuxue float64
	AvgYingyu float64
}

func (stu *Student) AvgScore() {
	stu.Avg = (stu.Yuwen + stu.Shuxue + stu.Yingyu) / 3
}

func (cls *Class) AvgScore() {
	total := len(cls.students)
	if total == 0 {
		return
	}
	sumYuwen := 0.0
	sumShuxue := 0.0
	sumYingyu := 0.0

	for _, stu := range cls.students {
		sumYingyu += stu.Yingyu
		sumShuxue += stu.Shuxue
		sumYuwen += stu.Yuwen
	}

	cls.AvgYuwen = sumYuwen / float64(total)
	cls.AvgShuxue = sumShuxue / float64(total)
	cls.AvgYingyu = sumYingyu / float64(total)
}

func (cls *Class) FailCount() int {
	rect := 0
	for _, stu := range cls.students {
		stu.AvgScore()
		if stu.Avg < 60 {
			rect++
		}
	}
	return rect
}
