package main

import "fmt"

func main() {
	x := 20
	fmt.Println("Value of a and memory location of a are: ", x, &x)
	fmt.Scan(&x)
	fmt.Println("Updated Value of a and memory location of a are: ", x, &x)
}
