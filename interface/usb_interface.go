package main

import "fmt"

type USB interface {
	start()
	end()
}

type Mouse struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标准备就绪")
}

func (m Mouse) end() {
	fmt.Println(m.name, "鼠标结束拔出")
}

type FlashDisk struct {
	name string
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "U盘准备就绪")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "U盘结束拔出")
}

func testInterface(usb USB) {
	usb.start()
	usb.end()
}

func main() {
	m := Mouse{"罗技"}
	fmt.Println(m.name)
	f := FlashDisk{"闪迪"}

	var usb USB
	usb = m
	usb = f

	usb.start()
	usb.end()

	fmt.Println("----------------------")
	testInterface(m)
	testInterface(f)
}
