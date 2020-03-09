package main

import (
	"fmt"
)

type Human interface {
	Walk()
}

type Animal interface {
	Walk()
}

type Man struct {
	name string
}

type Cat struct {
	name string
}

func MakeThemWalk(i interface{}) {
	fmt.Println(i.(Animal))
}

func main() {
	//h := Man{name: "Titus"}
	var h Cat = Cat{name: "Chutti"}
	MakeThemWalk(h)
}

func (m Man) Walk() {
	fmt.Println("Human walking")
}

func (c Cat) Walk() {
	fmt.Println("Cat walking")
}
