package main

import "fmt"

type Integer int

//func (a Integer) Add(b Integer) Integer {
//	return a + b
//}

func (a *Integer) Add(b Integer) {
	*a = (*a) + b
}

func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

//type Math interface {
//	Add(i Integer) Integer
//	Multiply(i Integer) Integer
//}

type Math interface {
	Add(i Integer)
	Multiply(i Integer) Integer
}

//func main() {
//	var a Integer = 1
//	var m Math = &a
//	fmt.Println(m.Add(1))
//}

func main() {
	var a, b Integer = 1, 2
	A := a
	var m Math = &a
	m.Add(b)
	fmt.Printf("%d + %d = %d\n", A, b, a)
}
