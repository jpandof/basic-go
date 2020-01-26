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

func main() {
	i := IronMan{Man{"Pando"}, Robot{78, "Jesus"}, "fuego"}
	fmt.Printf("SuperPoder, %v\n", i.SuperPoder)
	fmt.Printf("Name man, %v\n", i.Man.Name)
	fmt.Printf("Name robot, %v\n", i.Robot.Name)
	fmt.Printf("Batery, %v\n", i.Batery)
}
