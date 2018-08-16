package main

import "fmt"

func main() {
	func() {
		fmt.Println("Inside an anonymous self executing function")
	}()
}
