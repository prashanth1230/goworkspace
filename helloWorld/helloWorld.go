package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \t %b \t %#x \n", i, i, i)
	}
}
