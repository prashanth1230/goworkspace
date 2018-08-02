package main

import (
	Reverse "../reverse"
	"fmt"
)

// Package level variable
var x int = 0

func main() {
	fmt.Println(x)
	fmt.Println(Reverse.MyName)
	fmt.Println(Reverse.Add(1, 2))
}
