package main

import "fmt"

type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student{id: id, name: name, male: male, score: score}
}

//值方法用来查询，不会改变原结构体内容
func (s Student) GetName() string {
	return s.name
}

//指针方法用来修改原结构体内的内容
func (s *Student) SetName(name string) {
	s.name = name
}

func main() {
	student := NewStudent(1, "jiangsudaxue", false, 89)
	fmt.Println(student)
}
