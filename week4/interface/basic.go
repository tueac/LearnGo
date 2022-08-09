package main

type Transporter interface {
	move(string, string) (int, error)
	say(string)
}

type Car struct {
	Name  string
	Price int
}

func (Car) move(string, string) (int, error) {
	return 0, nil
}

func (Car) say(string) {

}

func (Car) hello() {

}

func main1() {
	var transp Transporter
	car := Car{"BMW", 100}
	transp = car //国为 Car 实现了接口 Transporter
	transp.move("a", "b")
	car.hello()
}
