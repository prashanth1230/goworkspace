package main

import (
	SwitchOnType "../switchOnType"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	name := "Prashanth"

	switch name {
	case "Prashanth":
		fmt.Println("Hello Prashanth")
	case "Shruthi":
		fmt.Println("Hello Shruthi")
	default:
		fmt.Println("Hello")
	}
	fallThrough(name)
	SwitchOnType.SwitchOnType(true)
}

func fallThrough(name string) {
	switch name {
	case "Prashanth":
		fmt.Println("Prashanth hi!")
		fallthrough
	case "Shruthi":
		fmt.Println("Shruthi hi!")
	default:
		fmt.Println("Hi!")
	}
	multipleSwitch(name)
}

func multipleSwitch(name string) {
	switch name {
	case "Prashanth", `Shruthi`:
		fmt.Println("Hello both")
	case "shashank":
		fmt.Println("Hello Shashank")
	}
}
