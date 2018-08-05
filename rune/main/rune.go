package main

import "fmt"

// Runes are just characters. UTF-8 character set is supported

func main() {
	fmt.Println('A')
	for i := 300; i < 310; i++ {
		fmt.Println(i, "String value of Rune ", string(i))
	}
}
