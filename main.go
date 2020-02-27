package main

import "fmt"

func walkTestOne(a Animal) (err error) {

	return nil
}

func walkTestTwo(a Animal) (err error) {

	return nil
}

func main() {
	human := Human{name: "Titus"}
	lion := Lion{name: "leo"}
	butterfly := Butterfly{name: "black mama"}

	var globalInterface Animal
	globalInterface = human
	globalInterface = lion

	walkTestOne(&human)
	walkTestTwo(&human)

	fmt.Println(human)
	fmt.Println(lion)
	fmt.Println(butterfly)
	fmt.Println(globalInterface)
}

type Animal interface {
	Walk() (err error)
}

type Human struct {
	name string
}

type Lion struct {
	name string
}

type Butterfly struct {
	name string
}

// Walk ia s property for human to walk
func (h Human) Walk() (err error) {
	fmt.Println("Human Walking")
	return nil
}

// Walk ia s property for Animal to walk
func (a Lion) Walk() (err error) {
	fmt.Println("Animal Walking")
	return nil
}
