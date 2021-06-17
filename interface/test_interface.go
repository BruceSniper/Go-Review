package main

import "fmt"

type dog struct{}
type cat struct{}

type person struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵")
}

func (d dog) say() {
	fmt.Println("汪汪")
}

func (p person) say() {
	fmt.Println("啊啊")
}

type Sayer interface {
	say()
}

func da(arg Sayer) {
	arg.say()
}

func main() {
	d := dog{}
	c := cat{}
	p := person{
		name: "ss",
	}
	d.say()
	da(d)
	c.say()
	da(c)
	p.say()
	da(p)
}
