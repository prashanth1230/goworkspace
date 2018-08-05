package main

import "fmt"

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
}
