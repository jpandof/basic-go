package main

import "fmt"

type Man struct {
	Name string
}

type Robot struct {
	Batery int
	Name   string
}

type IronMan struct {
	Man
	Robot
	SuperPoder string
}

func (i *IronMan) GetBatery() int {
	return i.Batery * 2
}

func (i *IronMan) getSuperPoder() string {
	return i.SuperPoder + "_fly"
}

func main() {
	i := IronMan{Man{"Pando"}, Robot{78, "Jesus"}, "fuego"}
	fmt.Printf("Batery, %v\n", i.GetBatery())
	fmt.Printf("Batery, %v\n", i.Batery)
	fmt.Printf("SuperPoder, %v\n", i.getSuperPoder())
	fmt.Printf("SuperPoder, %v\n", i.SuperPoder)
	fmt.Printf("Name man, %v\n", i.Man.Name)
	fmt.Printf("Name robot, %v\n", i.Robot.Name)

}
